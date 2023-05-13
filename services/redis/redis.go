package redis

import (
	"context"
	"fmt"
	boosterModels "github.com/InsuranceTech/shared/booster/model"
	"github.com/InsuranceTech/shared/common"
	"github.com/InsuranceTech/shared/common/depth"
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/config"
	"github.com/InsuranceTech/shared/log"
	"github.com/redis/go-redis/v9"
	"time"
)

// Msgpack : https://github.com/smallnest/gosercomp/blob/master/benchmark.png

var (
	Client *redis.Client
	_log   = log.CreateTag("Redis")
)

func Init(ctx context.Context, cfg *config.Config) {
	Client = redis.NewClient(&redis.Options{
		ClientName: cfg.Server.SERVICE_NAME,
		Addr:       fmt.Sprintf("%s:%d", cfg.Redis.HOST, cfg.Redis.PORT),
		Password:   cfg.Redis.PASS,
		DB:         cfg.Redis.DEFAULT_DB,
		OnConnect:  OnConnect,
	})
	status := Client.Ping(ctx)
	if status.Err() == nil {
		_log.Info("Connected")
	} else {
		_log.Fatal("Connection Error", status.Err())
	}
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

// GetkeyIndicatorResult
// Redisteki adresini döndürür
// Örnek: BINANCE_SPOT:BTCUSDT:5:Candles
func GetkeyIndicatorResult(symbol *symbol.Symbol, indicatorId int64) string {
	return fmt.Sprintf("%s:Indicators:%d", symbol.ToString(), indicatorId)
}

func GetkeyDepth(symbol *symbol.Symbol) string {
	return fmt.Sprintf("%s:Depth", symbol.ToStringNoPeriod())
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

func SaveIndicatorResult(data *boosterModels.IndicatorResult) error {
	key := GetkeyIndicatorResult(data.Symbol, data.IndicatorID)
	bytes, err := data.MarshalMsg(nil)
	if err != nil {
		_log.Error("SaveIndicatorResult.MarshalMsg", err, data.Symbol.ToString())
		return err
	}
	cmdStatus := Client.Set(context.Background(), key, bytes, 0)
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
	bytes, err := collection.MarshalMsg(nil)
	if err != nil {
		_log.Error("SaveIndicatorResultCollection.MarshalMsg", err)
		return err
	}
	cmdStatus := Client.Set(context.Background(), key, bytes, 0)
	if cmdStatus.Err() != nil {
		_log.Error("SaveIndicatorResultCollection.Redis.Set", cmdStatus.Err())
		return cmdStatus.Err()
	}

	cmdStatus = Client.Set(context.Background(), keyTime, updateTime.UnixMilli(), 0)
	if cmdStatus.Err() != nil {
		_log.Error("SaveIndicatorResultCollection.Redis.Set.Time", cmdStatus.Err())
		return cmdStatus.Err()
	}
	return nil
}

func GetIndicatorResultCollection() (*boosterModels.IndicatorResultCollection, error) {
	data := &boosterModels.IndicatorResultCollection{}
	err := UpdateIndicatorResultCollectionModel(data)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func UpdateIndicatorResultCollectionModel(data *boosterModels.IndicatorResultCollection) error {
	key := "INDICATOR_RESULTS"
	cmdStatus := Client.Get(context.Background(), key)
	if cmdStatus.Err() != nil {
		return cmdStatus.Err()
	}
	bytes, err := cmdStatus.Bytes()
	if err != nil {
		return err
	}
	_, err = data.UnmarshalMsg(bytes)
	if err != nil {
		_log.Error("UpdateIndicatorResultCollectionModel.UnmarshalMsg", err)
		return err
	}
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
