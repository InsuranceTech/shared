package nats

import (
	"fmt"
	"github.com/InsuranceTech/shared/common"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/config"
	"github.com/nats-io/nats.go"
)

var Client *nats.Conn

func Init(cfg *config.Config) {
	connStr := fmt.Sprintf("nats://%s:%s@%s:%d", cfg.Nats.USER, cfg.Nats.PASS, cfg.Nats.HOST, cfg.Nats.CLIENT_PORT)
	println(connStr)
	err := error(nil)
	option := func(option *nats.Options) error {
		option.ClosedCB = onCloseConnection
		option.DisconnectedErrCB = onDisconnectConnection
		option.ConnectedCB = onConnected
		option.ReconnectedCB = onReConnected
		return nil
	}
	Client, err = nats.Connect(connStr, option)
	if err != nil {
		panic(err)
	}
}

func onReConnected(conn *nats.Conn) {
	fmt.Println("NATSIO : ", "Reconnected", conn.Opts.Url)
}

func onConnected(conn *nats.Conn) {
	fmt.Println("NATSIO : ", "Connected", conn.Opts.Url)
}

func onCloseConnection(conn *nats.Conn) {
	fmt.Println("NATSIO : ", "Closed Connection", conn.Opts.Url)
}

func onDisconnectConnection(conn *nats.Conn, err error) {
	fmt.Println("NATSIO : ", "Disconnected", conn.Opts.Url, err)
}

func OnClosedCandleSymbol(symbol *symbol.Symbol, handler func(candle *common.Candle)) {
	subject := fmt.Sprintf("%s.ClosedCandle", symbol.ToString())
	Client.Subscribe(subject, func(msg *nats.Msg) {
		closedSymbol := msg.Header.Get("symbol")
		fmt.Println(closedSymbol)
		candle := &common.Candle{}
		_, err := candle.UnmarshalMsg(msg.Data)
		if err != nil {
			panic(err)
		}
		go func() {
			handler(candle)
		}()
	})
}

func OnClosedCandleSymbols(handler func(symbol *symbol.Symbol, candle *common.Candle)) {
	subject := "*.ClosedCandle"
	Client.Subscribe(subject, func(msg *nats.Msg) {
		closedSymbol := msg.Header.Get("symbol")
		symbol, parseOk := symbol.ParseSymbolEx(closedSymbol)
		if parseOk == false {
			panic("Symbol parse error : " + closedSymbol)
		}
		candle := &common.Candle{}
		_, err := candle.UnmarshalMsg(msg.Data)
		if err != nil {
			panic(err)
		}
		go func() {
			handler(symbol, candle)
		}()
	})
}
