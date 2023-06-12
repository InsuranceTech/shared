package nats

import (
	"encoding/json"
	"fmt"
	boosterModels "github.com/InsuranceTech/shared/booster/model"
	"github.com/InsuranceTech/shared/common"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/config"
	"github.com/InsuranceTech/shared/log"
	model3 "github.com/InsuranceTech/shared/services/nats/model"
	"github.com/nats-io/nats.go"
	"github.com/spf13/cast"
	"strconv"
)

const (
	SUBJECT_KEY_ChangedIndicatorCollection = "CHANGED_IDNICATOR_COLLECTION"
	SUBJECT_KEY_INDICATOR_SCANNER          = "INDICATOR_SCANNER"
)

var (
	Client *nats.Conn
	_log   = log.CreateTag("Nats")
)

func Init(cfg *config.Config) {
	connStr := fmt.Sprintf("nats://%s:%s@%s:%d", cfg.Nats.USER, cfg.Nats.PASS, cfg.Nats.HOST, cfg.Nats.CLIENT_PORT)
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
		_log.Fatal("Nats bağlantısı kurulamadı.", err)
	}
}

func onReConnected(conn *nats.Conn) {
	_log.Info("Reconnected ", conn.Opts.Url)
}

func onConnected(conn *nats.Conn) {
	_log.Info("Connected ", conn.Opts.Url)
}

func onCloseConnection(conn *nats.Conn) {
	_log.Info("Closed Connection ", conn.Opts.Url)
}

func onDisconnectConnection(conn *nats.Conn, err error) {
	if err == nil {
		_log.Info("Disconnect ", conn.Opts.Url)
	} else {
		_log.Error("Disconnect ", conn.Opts.Url, err)
	}
}

// region Subscribes
func OnClosedCandleSymbol(symbol *symbol.Symbol, handler func(candle *common.Candle)) {
	subject := fmt.Sprintf("%s.ClosedCandle", symbol.ToString())
	Client.Subscribe(subject, func(msg *nats.Msg) {
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
		_symbol := msg.Header.Get("symbol")
		symbol, parseOk := symbol.ParseSymbolEx(_symbol)
		if parseOk == false {
			_log.Error("OnClosedCandleSymbols", "Symbol parse error : "+_symbol)
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

func OnChangeBoosterSignals(handler func(symbol *symbol.Symbol, funcName string, signal int16, indicatorId int64)) {
	subject := "*.BoosterSignal"
	Client.Subscribe(subject, func(msg *nats.Msg) {
		_symbol := msg.Header.Get("symbol")
		symbol, parseOk := symbol.ParseSymbolEx(_symbol)
		if parseOk == false {
			_log.Error("Nats.OnChangeBoosterSignals", "Symbol parse error : "+_symbol)
			return
		}
		_funcName := msg.Header.Get("funcName")
		_indicatorId := msg.Header.Get("indicatorId")
		_signal := msg.Header.Get("signal")
		indicatorId, err := cast.ToInt64E(_indicatorId)
		if err != nil {
			_log.Error("Nats.OnChangeBoosterSignals", "indicatorId casting error", err)
			return
		}
		signal, err := cast.ToInt16E(_signal)
		if err != nil {
			_log.Error("Nats.OnChangeBoosterSignals", "signal casting error", err)
			return
		}
		go func() {
			handler(symbol, _funcName, signal, indicatorId)
		}()
	})
}

// OnChangeIndicatorCollection Redisteki gösterge koleksiyonu değiştirildiğinde tetiklenir
func OnChangeIndicatorCollection(handler func()) (*nats.Subscription, error) {
	subject := SUBJECT_KEY_ChangedIndicatorCollection
	return Client.Subscribe(subject, func(msg *nats.Msg) {
		go func() {
			handler()
		}()
	})
}

func OnRequestIndicatorScanner(handler func(request *model3.ReqIndicatorScanner) model3.ReqIndicatorScanner) (*nats.Subscription, error) {
	subject := SUBJECT_KEY_INDICATOR_SCANNER
	return Client.Subscribe(subject, func(msg *nats.Msg) {
		var req *model3.ReqIndicatorScanner
		err := json.Unmarshal(msg.Data, &req)

		// Error JSON Deserialize
		if err != nil {
			response := model3.ResIndicatorScanner{
				ResultCode: -1, // Error json format
				Results:    nil,
			}
			responseBytes, _ := json.Marshal(response)
			msg.Respond(responseBytes)
			return
		}

		handlerResponse := handler(req)
		responseBytes, _ := json.Marshal(handlerResponse)
		msg.Respond(responseBytes)
	})
}

//endregion

// region Triggers
func TriggerClosedCandle(symbol *symbol.Symbol, candle *common.Candle) {
	candleBytes, err := candle.MarshalMsg(nil)
	if err != nil {
		_log.Error("TriggerClosedCandle", "Candle.MarshalMsg", err)
		return
	}
	msg := nats.Msg{
		Subject: fmt.Sprintf("%s.ClosedCandle", symbol.ToString()),
		Header:  map[string][]string{"symbol": {symbol.ToString()}},
		Data:    candleBytes,
	}
	err = Client.PublishMsg(&msg)
	if err != nil {
		_log.Error("TriggerClosedCandle", "PublishMsg", err)
		return
	}
}

func TriggerChangedBoosterSignal(data *boosterModels.IndicatorResult) {
	msg := nats.Msg{
		Subject: fmt.Sprintf("%s.%s.BoosterSignal", data.Symbol.ToString(), data.FuncName),
		Header: map[string][]string{
			"symbol":      {data.Symbol.ToString()},
			"funcName":    {data.FuncName},
			"indicatorId": {strconv.Itoa(int(data.IndicatorID))},
			"signal":      {strconv.Itoa(int(data.Signal))},
		},
		Data: nil,
	}
	err := Client.PublishMsg(&msg)
	if err != nil {
		_log.Error("TriggerChangedBoosterSignal", "PublishMsg", err)
	}
}

// TriggerChangedIndicatorCollection Göstergeler hesaplandıktan sonra toplu halde redise kaydedikten sonra tetiklenir
func TriggerChangedIndicatorCollection() {
	msg := nats.Msg{
		Subject: SUBJECT_KEY_ChangedIndicatorCollection,
	}
	err := Client.PublishMsg(&msg)
	if err != nil {
		_log.Error("TriggerChangedIndicatorCollection", "PublishMsg", err)
	}
}

//endregion
