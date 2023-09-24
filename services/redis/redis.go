package redis

import (
	"context"
	"encoding/json"
	"fmt"
	boosterModels "github.com/InsuranceTech/shared/booster/model"
	"github.com/InsuranceTech/shared/common"
	"github.com/InsuranceTech/shared/common/depth"
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/ratio"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/config"
	"github.com/InsuranceTech/shared/log"
	"github.com/InsuranceTech/shared/services/redis/model"
	sqlModel "github.com/InsuranceTech/shared/services/sql/model"
	"github.com/korovkin/limiter"
	"github.com/redis/go-redis/v9"
	"strconv"
	"sync"
	"time"
)

// Msgpack : https://github.com/smallnest/gosercomp/blob/master/benchmark.png

const (
	_ALARM_INDICATOR_PREFIX = "ALARM:INDICATOR"
	_ALARM_BOOSTER_PREFIX   = "ALARM:BOOSTER"
)

var (
	Client       *redis.Client
	_ctx         context.Context
	_log         = log.CreateTag("Redis")
	_limitWorker *limiter.ConcurrencyLimiter
)

func Init(ctx context.Context, cfg *config.Config) {
	Client = redis.NewClient(&redis.Options{
		ClientName: cfg.Server.SERVICE_NAME,
		Addr:       fmt.Sprintf("%s:%d", cfg.Redis.HOST, cfg.Redis.PORT),
		Password:   cfg.Redis.PASS,
		DB:         cfg.Redis.DEFAULT_DB,
		OnConnect:  OnConnect,
	})
	_ctx = ctx
	status := Client.Ping(ctx)
	if status.Err() == nil {
		_log.Info("Connected")
	} else {
		_log.Fatal("Connection Error", status.Err())
	}
}

func SetLimitWorker(limit int) {
	_limitWorker = limiter.NewConcurrencyLimiter(limit)
}

func OnConnect(ctx context.Context, conn *redis.Conn) error {
	_log.Info("Connected")
	return nil
}

// GetkeyCandles
// Redisteki adresini döndürür
// Örnek: BINANCE_SPOT:BTCUSDT:5:Candles
func GetkeyCandles(symbol *symbol.Symbol) string {
	return fmt.Sprintf("%s:Candles", symbol.ToString())
}

// GetkeyOpenCandle
// Redisteki adresini döndürür
// Örnek: BINANCE_SPOT:BTCUSDT:5:Open
func GetkeyOpenCandle(symbol *symbol.Symbol) string {
	return fmt.Sprintf("%s:Open", symbol.ToString())
}

// GetkeyIndicatorResult
// Redisteki adresini döndürür
// Örnek: BINANCE_SPOT:BTCUSDT:5:Candles
func GetkeyIndicatorResult(symbol *symbol.Symbol, indicatorId int64) string {
	return fmt.Sprintf("%s:Indicators:%d", symbol.ToString(), indicatorId)
}

func GetkeyDepth(symbol *symbol.Symbol) string {
	return fmt.Sprintf("%s:Depth", symbol.ToStringNoPeriod())
}

func GetkeyLongShortRatio(symbol *symbol.Symbol) string {
	return fmt.Sprintf("%s:TLSR", symbol.ToString())
}

func GetkeyTick(symbol *symbol.Symbol) string {
	return fmt.Sprintf("%s:Tick", symbol.ToStringNoPeriod())
}

func SaveCandleSeries(periodicSeries *common.PeriodicCandleSeries, period period.Period) bool {
	series := periodicSeries.Get(period)
	if series == nil {
		return false
	}
	bytes, err := series.MarshalMsg(nil)
	if err != nil {
		_log.Error("SaveCandleSeries.MarshalMsg", err)
		return false
	}

	key := GetkeyCandles(periodicSeries.Symbol.CloneP(period))
	status := Client.Set(context.Background(), key, bytes, 0)
	if status.Err() != nil {
		_log.Error("SaveCandleSeries.Redis.Set", status.Err())
		return false
	}
	return true
}

func SaveCandleSeriesAsync(periodicSeries *common.PeriodicCandleSeries, period period.Period, onComplete func(status bool)) {
	go func() {
		status := SaveCandleSeries(periodicSeries, period)
		if onComplete != nil {
			onComplete(status)
		}
	}()
}

func SaveCandleSeriesLWAsync(periodicSeries *common.PeriodicCandleSeries, period period.Period, onComplete func(status bool)) {
	_limitWorker.Execute(func() {
		status := SaveCandleSeries(periodicSeries, period)
		if onComplete != nil {
			onComplete(status)
		}
	})
}

func GetCandleSeries(symbol *symbol.Symbol) (*common.CandleSeries, error) {
	key := GetkeyCandles(symbol)
	cmdStatus := Client.Get(context.Background(), key)
	if cmdStatus.Err() != nil {
		return nil, cmdStatus.Err()
	}
	bytes, err := cmdStatus.Bytes()
	if err != nil {
		return nil, err
	}
	series := &common.CandleSeries{}
	_, err = series.UnmarshalMsg(bytes)
	if err != nil {
		_log.Error("GetCandleSeries.UnMarshalMsg", err, symbol.ToString())
		return nil, err
	}
	return series, nil
}

func SaveDepthData(symbol *symbol.Symbol, depthData *depth.DepthData) error {
	key := GetkeyDepth(symbol)
	bytes, err := depthData.MarshalMsg(nil)
	if err != nil {
		_log.Error("SaveDepthData.MarshalMsg", err, symbol.ToString())
		return err
	}
	cmdStatus := Client.Set(context.Background(), key, bytes, 0)
	if cmdStatus.Err() != nil {
		_log.Error("SaveDepthData.Redis.Set", err, symbol.ToString())
		return cmdStatus.Err()
	}
	return nil
}

func GetDepthData(symbol *symbol.Symbol) (*depth.DepthData, error) {
	redisKey := GetkeyDepth(symbol)
	cmdStatus := Client.Get(context.Background(), redisKey)
	if cmdStatus.Err() != nil {
		return nil, cmdStatus.Err()
	}
	bytes, _ := cmdStatus.Bytes()
	depthData := &depth.DepthData{}
	_, err := depthData.UnmarshalMsg(bytes)
	if err != nil {
		_log.Error("GetDepthData.UnmarshalMsg", err, symbol.ToString())
		return nil, err
	}
	return depthData, nil
}

func SaveLongShortRatioData(symbol *symbol.Symbol, data *ratio.TakeLongData) error {
	key := GetkeyLongShortRatio(symbol)
	bytes, err := json.Marshal(data)
	if err != nil {
		_log.Error("SaveLongShortRatioData.MarshalMsg", err, symbol.ToString())
		return err
	}
	cmdStatus := Client.Set(context.Background(), key, bytes, 0)
	if cmdStatus.Err() != nil {
		_log.Error("SaveLongShortRatioData.Redis.Set", err, symbol.ToString())
		return cmdStatus.Err()
	}
	return nil
}

func GetLongShortRatioData(symbol *symbol.Symbol) (*ratio.TakeLongData, error) {
	redisKey := GetkeyLongShortRatio(symbol)
	cmdStatus := Client.Get(context.Background(), redisKey)
	if cmdStatus.Err() != nil {
		return nil, cmdStatus.Err()
	}
	bytes, _ := cmdStatus.Bytes()
	data := &ratio.TakeLongData{}
	err := json.Unmarshal(bytes, data)
	if err != nil {
		_log.Error("GetLongShortRatioData.UnmarshalMsg", err, symbol.ToString())
		return nil, err
	}
	return data, nil
}

func SaveIndicatorResult(data *boosterModels.IndicatorResult) error {
	return saveIndicatorResultTtl(data, 0)
}

func SaveIndicatorResultTTL(data *boosterModels.IndicatorResult) error {
	return saveIndicatorResultTtl(data, data.Symbol.Period.Duration()+(time.Second*30))
}

func saveIndicatorResultTtl(data *boosterModels.IndicatorResult, ttl time.Duration) error {
	key := GetkeyIndicatorResult(data.Symbol, data.IndicatorID)
	bytes, err := json.Marshal(data)
	if err != nil {
		_log.Error("SaveIndicatorResult.MarshalMsg", err, data.Symbol.ToString())
		return err
	}
	cmdStatus := Client.Set(context.Background(), key, bytes, ttl)
	if cmdStatus.Err() != nil {
		_log.Error("SaveIndicatorResult.Redis.Set", cmdStatus.Err(), data.Symbol.ToString())
		return cmdStatus.Err()
	}
	return nil
}

func GetIndicatorResult(symbol *symbol.Symbol, indicatorID int64) (*boosterModels.IndicatorResult, error) {
	redisKey := GetkeyIndicatorResult(symbol, indicatorID)
	cmdStatus := Client.Get(context.Background(), redisKey)
	if cmdStatus.Err() != nil {
		return nil, cmdStatus.Err()
	}
	bytes, _ := cmdStatus.Bytes()
	data := &boosterModels.IndicatorResult{}
	_, err := data.UnmarshalMsg(bytes)
	if err != nil {
		_log.Error("GetIndicatorResult.UnmarshalMsg", err, data.Symbol.ToString())
		return nil, err
	}
	return data, nil
}

func SaveIndicatorResultCollection(collection *boosterModels.IndicatorResultCollection) error {
	updateTime := time.Now().UTC()
	key := "INDICATOR_RESULTS"
	keyTime := "INDICATOR_RESULTS.TIME"

	wait := sync.WaitGroup{}
	for keyPeriod, _ := range collection.Results {
		wait.Add(1)
		go func(p period.Period) {
			bytes := make([]byte, 0)
			for _, result := range collection.GetIndicatorsP(p) {
				bytes, _ = result.MarshalMsg(bytes)
			}
			cmdStatus := Client.Set(context.Background(), key+"_"+p.MinuteStr(), bytes, 0)
			if cmdStatus.Err() != nil {
				_log.Error("SaveIndicatorResultCollection.Redis.Set", cmdStatus.Err())
			}
			wait.Done()
		}(keyPeriod)
	}
	wait.Wait()
	cmdStatus := Client.Set(context.Background(), keyTime, updateTime.UnixMilli(), 0)
	if cmdStatus.Err() != nil {
		_log.Error("SaveIndicatorResultCollection.Redis.Set.Time", cmdStatus.Err())
		return cmdStatus.Err()
	}
	return nil
}

func GetIndicatorResultCollection() (*boosterModels.IndicatorResultCollection, error) {
	data := boosterModels.CreateIndicatorResultCollection()
	err := UpdateIndicatorResultCollectionModel(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateIndicatorResultCollectionModel(data *boosterModels.IndicatorResultCollection) error {
	key := "INDICATOR_RESULTS"
	keyTime := "INDICATOR_RESULTS.TIME"
	lastTimeCmd := Client.Get(context.Background(), keyTime)
	if lastTimeCmd.Err() == nil {
		data.LastTime, _ = lastTimeCmd.Int64()
	}

	wait := sync.WaitGroup{}
	for _, keyPeriod := range period.AllPeriods {
		wait.Add(1)
		go func(p period.Period) {
			cmdStatus := Client.Get(context.Background(), key+"_"+p.MinuteStr())
			if cmdStatus.Err() != nil {
				wait.Done()
				return
			}
			bytes, err := cmdStatus.Bytes()
			if err != nil {
				wait.Done()
				return
			}

			for len(bytes) > 0 {
				r := &boosterModels.IndicatorResult{}
				bytes, _ = r.UnmarshalMsg(bytes)
				data.Append(r)
			}
			wait.Done()
		}(keyPeriod)
	}
	wait.Wait()
	data.Indexes()
	return nil
}

func GetIndicatorResultCollectionTime() (int64, error) {
	key := "INDICATOR_RESULTS.TIME"
	cmdStatus := Client.Get(context.Background(), key)
	if cmdStatus.Err() != nil {
		return 0, cmdStatus.Err()
	}
	unix, err := cmdStatus.Int64()
	if err != nil {
		return 0, err
	}
	return unix, nil
}

func UpdateIndicatorResultCollectionModelIfNeed(c *boosterModels.IndicatorResultCollection) (updated bool, err error) {
	updateTime, err := GetIndicatorResultCollectionTime()
	if err != nil {
		return false, err
	}
	if c.LastTime != updateTime {
		return true, UpdateIndicatorResultCollectionModel(c)
	}
	return false, nil
}

func SaveTickData(symbol *symbol.Symbol, data *model.BaseTickData) error {
	key := GetkeyTick(symbol)
	bytes, err := data.MarshalMsg(nil)
	if err != nil {
		_log.Error("SaveTickData.MarshalMsg", err, symbol.ToString())
		return err
	}
	cmdStatus := Client.Set(context.Background(), key, bytes, 0)
	if cmdStatus.Err() != nil {
		_log.Error("SaveTickData.Redis.Set", err, symbol.ToString())
		return cmdStatus.Err()
	}
	return nil
}

func GetTickData(symbol *symbol.Symbol) (*model.BaseTickData, error) {
	redisKey := GetkeyTick(symbol)
	cmdStatus := Client.Get(context.Background(), redisKey)
	if cmdStatus.Err() != nil {
		return nil, cmdStatus.Err()
	}
	bytes, _ := cmdStatus.Bytes()
	data := &model.BaseTickData{}
	_, err := data.UnmarshalMsg(bytes)
	if err != nil {
		_log.Error("GetTickData.UnmarshalMsg", err, symbol.ToString())
		return nil, err
	}
	return data, nil
}

func GetAllIndicatorAlarms() ([]*sqlModel.AlarmIndicator, error) {
	scanCmd := Client.Keys(_ctx, _ALARM_INDICATOR_PREFIX+":*")
	if scanCmd.Err() != nil {
		return nil, scanCmd.Err()
	}
	keys, _ := scanCmd.Result()

	valuesCmd := Client.MGet(_ctx, keys...)
	if valuesCmd.Err() != nil {
		return nil, valuesCmd.Err()
	}

	resultStrings, _ := valuesCmd.Result()
	results := make([]*sqlModel.AlarmIndicator, len(resultStrings))
	for i, resultString := range resultStrings {
		data := &sqlModel.AlarmIndicator{}
		_ = json.Unmarshal([]byte(resultString.(string)), &data)
		results[i] = data
	}

	return results, nil
}

func SetAllIndicatorAlarms(data []*sqlModel.AlarmIndicator) (bool, error) {
	var items []interface{}
	for _, i := range data {
		bytes, _ := json.Marshal(i)
		items = append(items, _ALARM_INDICATOR_PREFIX+":"+strconv.Itoa(i.ID), bytes)
	}
	status := Client.MSet(_ctx, items)
	if status.Err() != nil {
		return false, status.Err()
	} else {
		return true, nil
	}
}

func ClearAllIndicatorAlarms() (bool, error) {
	scanCmd := Client.Keys(_ctx, _ALARM_INDICATOR_PREFIX+":*")
	if scanCmd.Err() != nil {
		return false, scanCmd.Err()
	}
	keys, _ := scanCmd.Result()

	cmd := Client.Del(_ctx, keys...)

	if cmd.Err() != nil {
		return false, cmd.Err()
	}

	return true, nil
}

func GetAllBoosterAlarms() ([]*sqlModel.AlarmBooster, error) {
	scanCmd := Client.Keys(_ctx, _ALARM_INDICATOR_PREFIX+":*")
	if scanCmd.Err() != nil {
		return nil, scanCmd.Err()
	}
	keys, _ := scanCmd.Result()

	valuesCmd := Client.MGet(_ctx, keys...)
	if valuesCmd.Err() != nil {
		return nil, valuesCmd.Err()
	}

	resultStrings, _ := valuesCmd.Result()
	results := make([]*sqlModel.AlarmBooster, len(resultStrings))
	for i, resultString := range resultStrings {
		data := &sqlModel.AlarmBooster{}
		_ = json.Unmarshal([]byte(resultString.(string)), &data)
		results[i] = data
	}

	return results, nil
}

func SetAllBoosterAlarms(data []*sqlModel.AlarmBooster) (bool, error) {
	var items []interface{}
	for _, i := range data {
		bytes, _ := json.Marshal(i)
		items = append(items, _ALARM_INDICATOR_PREFIX+":"+strconv.Itoa(i.ID), bytes)
	}
	status := Client.MSet(_ctx, items)
	if status.Err() != nil {
		return false, status.Err()
	} else {
		return true, nil
	}
}

func ClearAllBoosterAlarms() (bool, error) {
	scanCmd := Client.Keys(_ctx, _ALARM_INDICATOR_PREFIX+":*")
	if scanCmd.Err() != nil {
		return false, scanCmd.Err()
	}
	keys, _ := scanCmd.Result()

	cmd := Client.Del(_ctx, keys...)

	if cmd.Err() != nil {
		return false, cmd.Err()
	}

	return true, nil
}

func GetOpenCandle(symbol *symbol.Symbol) (*common.Candle, error) {
	key := GetkeyOpenCandle(symbol)
	cmdStatus := Client.Get(context.Background(), key)
	if cmdStatus.Err() != nil {
		return nil, cmdStatus.Err()
	}
	bytes, err := cmdStatus.Bytes()
	if err != nil {
		return nil, err
	}
	c := &common.Candle{}
	_, err = c.UnmarshalMsg(bytes)
	if err != nil {
		_log.Error("GetOpenCandle.UnMarshalMsg", err, symbol.ToString())
		return nil, err
	}
	return c, nil
}

func SaveOpenCandle(symbol *symbol.Symbol, candle *common.Candle) bool {
	if candle == nil {
		return false
	}
	bytes, err := candle.MarshalMsg(nil)
	if err != nil {
		_log.Error("SaveOpenCandle.MarshalMsg", err)
		return false
	}

	key := GetkeyOpenCandle(symbol)
	status := Client.Set(context.Background(), key, bytes, 0)
	if status.Err() != nil {
		_log.Error("SaveOpenCandle.Redis.Set", status.Err())
		return false
	}
	return true
}
