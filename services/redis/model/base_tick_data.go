package model

//msgp:tuple BaseTickData
//go:generate msgp -o=base_tick_data_msgpack.go -tests=false

type BaseTickData struct {
	PriceChange        float64
	PriceChangePercent float64
	LastPrice          float64
	CloseQty           float64
	BidPrice           float64
	BidQty             float64
	AskPrice           float64
	AskQty             float64
	OpenPrice          float64
	HighPrice          float64
	LowPrice           float64
	BaseVolume         float64
	QuoteVolume        float64
	OpenTime           int64
	CloseTime          int64
}
