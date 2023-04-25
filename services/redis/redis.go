package redis

import (
	"context"
	"fmt"
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
