package redis

import (
	"context"
	"fmt"
	"github.com/InsuranceTech/shared/common"
	"github.com/InsuranceTech/shared/common/depth"
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/config"
	"github.com/redis/go-redis/v9"
)

// Msgpack : https://github.com/smallnest/gosercomp/blob/master/benchmark.png

var (
	Client *redis.Client
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
		fmt.Println("Redis : ", "Connected")
	} else {
		fmt.Println("Redis : ", "Connection error!")
		panic(status.Err())
	}
}

func OnConnect(ctx context.Context, conn *redis.Conn) error {
	fmt.Println("Redis : ", "Connected")
	return nil
}

// GetkeyCandles
// Redisteki adresini döndürür
// Örnek: BINANCE_SPOT:BTCUSDT:5:Candles
func GetkeyCandles(symbol *symbol.Symbol) string {
	return fmt.Sprintf("%s:Candles", symbol.ToString())
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
		panic(err)
	}

	key := GetkeyCandles(periodicSeries.Symbol.CloneP(period))
	status := Client.Set(context.Background(), key, bytes, 0)
	if status.Err() != nil {
		panic(status.Err())
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
		return nil, err
	}
	return series, nil
}

func SaveDepthData(symbol *symbol.Symbol, depthData *depth.DepthData) error {
	key := GetkeyDepth(symbol)
	bytes, err := depthData.MarshalMsg(nil)
	if err != nil {
		return err
	}
	cmdStatus := Client.Set(context.Background(), key, bytes, 0)
	if cmdStatus.Err() != nil {
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
		return nil, err
	}
	return depthData, nil
}
