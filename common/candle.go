//msgp:tuple Candle
//go:generate msgp -o=Candle_Msgpack.go -tests=false
package common

import (
	"time"
)

type Candle struct {
	Date             *time.Time
	Open             float64
	High             float64
	Low              float64
	Close            float64
	Volume           float64
	QuoteAssetVolume float64
	TakerBaseVolume  float64
	TakerQuoteVolume float64
}

func NewCandleF(t *time.Time, value float64) *Candle {
	return &Candle{
		Date:             t,
		Open:             value,
		High:             value,
		Low:              value,
		Close:            value,
		Volume:           value,
		QuoteAssetVolume: value,
		TakerBaseVolume:  value,
		TakerQuoteVolume: value,
	}
}

func NewCandleFP(t time.Time, value float64) *Candle {
	return &Candle{
		Date:             &t,
		Open:             value,
		High:             value,
		Low:              value,
		Close:            value,
		Volume:           value,
		QuoteAssetVolume: value,
		TakerBaseVolume:  value,
		TakerQuoteVolume: value,
	}
}

func (candle *Candle) GetSource(_type SourceType) float64 {
	if _type == ST_OPEN {
		return candle.Open
	} else if _type == ST_HIGH {
		return candle.High
	} else if _type == ST_LOW {
		return candle.Low
	} else if _type == ST_CLOSE {
		return candle.Close
	} else if _type == ST_VOLUME {
		return candle.Volume
	} else if _type == ST_HL2 {
		return (candle.High + candle.Low) / 2
	} else if _type == ST_HLC3 {
		return (candle.High + candle.Low + candle.Close) / 3
	} else if _type == ST_HLCC4 {
		return (candle.High + candle.Low + candle.Close + candle.Close) / 4
	} else if _type == ST_OHLCV5 {
		return (candle.Open + candle.High + candle.Low + candle.Close + candle.Volume) / 5
	} else if _type == ST_OHLC4 {
		return (candle.Open + candle.High + candle.Low + candle.Close) / 4
	} else {
		panic("Unknown SourceType !")
	}
}

// GetSourcePtr Değişkenin adresini dönmek için yukarıdaki fonksiyonu tekrarladım
// eğer değişkene atayıp onun pointerini dönseydik yeni bir pointer tutan değişken oluşacaktı
// çirkin görünsede, performans için tercih edillmeli - (hesaplama yapılan tarafta mecburi - hl2,hlc3..)
func (candle *Candle) GetSourcePtr(_type SourceType) *float64 {
	if _type == ST_OPEN {
		return &candle.Open
	} else if _type == ST_HIGH {
		return &candle.High
	} else if _type == ST_LOW {
		return &candle.Low
	} else if _type == ST_CLOSE {
		return &candle.Close
	} else if _type == ST_VOLUME {
		return &candle.Volume
	} else if _type == ST_HL2 {
		ret := (candle.High + candle.Low) / 2
		return &ret
	} else if _type == ST_HLC3 {
		ret := (candle.High + candle.Low + candle.Close) / 3
		return &ret
	} else if _type == ST_HLCC4 {
		ret := (candle.High + candle.Low + candle.Close + candle.Close) / 4
		return &ret
	} else if _type == ST_OHLCV5 {
		ret := (candle.Open + candle.High + candle.Low + candle.Close + candle.Volume) / 5
		return &ret
	} else if _type == ST_OHLC4 {
		ret := (candle.Open + candle.High + candle.Low + candle.Close) / 4
		return &ret
	} else {
		panic("Unknown SourceType !")
	}
}
