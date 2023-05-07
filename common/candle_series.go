//go:generate msgp -o=CandleSeries_Msgpack.go -tests=false
package common

import (
	"errors"
	"github.com/InsuranceTech/shared/common/period"
	"math"
	"sort"
	"time"
)

type CandleSeries struct {
	Candles        []*Candle                                                `msg:"Candles"`
	MaxCount       int                                                      `msg:"MaxCount"`
	OnChangeSeries func(series *CandleSeries, candle *Candle, isAdded bool) `json:"-",msg:"-"`
}

// AddCandleE Son mum tarihini kontrol eder, aynıysa günceller
func (series *CandleSeries) AddCandleE(candle *Candle) (bool, error) {
	candleLen := len(series.Candles)

	if candleLen > 0 && candle.Date.Unix() < series.Candles[len(series.Candles)-1].Date.Unix() {
		// Hata
		return false, errors.New("Son mumdan daha eski tarihli bir mum eklenmeye çalıştı.")
	} else if candleLen > 0 && series.Candles[len(series.Candles)-1].Date.Unix() == candle.Date.Unix() {
		// Güncelle
		series.Candles[len(series.Candles)-1] = candle
		if series.OnChangeSeries != nil {
			series.OnChangeSeries(series, candle, false)
		}
		return false, nil
	} else {
		if series.MaxCount > 0 && candleLen >= series.MaxCount {
			// Kaydır
			series.Candles = series.Candles[1:candleLen]
		}
		// Ekle
		series.Candles = append(series.Candles, candle)
		if series.OnChangeSeries != nil {
			series.OnChangeSeries(series, candle, true)
		}
		return true, nil
	}
}

// AddCandle Son mum tarihini kontrol eder, aynıysa günceller
func (series *CandleSeries) AddCandle(candle *Candle) bool {
	s, _ := series.AddCandleE(candle)
	return s
}

// AddCandles Son mum tarihini kontrol eder, aynıysa günceller
func (series *CandleSeries) AddCandles(candles *[]*Candle) {
	for i := 0; i < len(*candles); i++ {
		series.AddCandle((*candles)[i])
	}
}

func (series *CandleSeries) Len() int {
	return len(series.Candles)
}

// InsertCandle Son mum tarihinin kontrolünü yapmadan ekler
func (series *CandleSeries) InsertCandle(candle *Candle) {
	len := len(series.Candles)
	if series.MaxCount > 0 && len >= series.MaxCount {
		series.Candles = series.Candles[1:len]
	}
	series.Candles = append(series.Candles, candle)
}

// InsertCandles Son mum tarihini kontrol eder, aynıysa günceller
func (series *CandleSeries) InsertCandles(candles *[]*Candle) {
	for i := 0; i < len(*candles); i++ {
		series.InsertCandle((*candles)[i])
	}
}

// Get Dikkat: index : 0 son elemanı getirir. 1 : bir önceki. (TradingView pinescriptteki gibi)
func (series *CandleSeries) Get(index int) *Candle {
	_index := len(series.Candles) - 1 - index
	if _index < 0 {
		return nil
	}
	return series.Candles[_index]
}

// TakeLastArr arr : Float64 çıktısı,
// count : seriesten fazlası istenildiğinde seriesten toplan alınan count/*
func (series *CandleSeries) TakeLastArr(takeCount int, sourceType SourceType) (arr *CandleArr, count int) {
	seriesLen := len(series.Candles)
	startIndex := seriesLen - takeCount
	if startIndex < 0 {
		ret := make(CandleArr, 0)
		return &ret, 0
	}
	ret := make(CandleArr, seriesLen-startIndex)
	for i := startIndex; i < seriesLen; i++ {
		ret[i-startIndex] = series.Candles[i].GetSourcePtr(sourceType)
	}
	return &ret, len(ret)
}

func (series *CandleSeries) MinDate() *time.Time {
	var t *time.Time = nil
	min := int64(math.MaxInt64)
	for i := 0; i < len(series.Candles); i++ {
		if series.Candles[i].Date.UnixMilli() < min {
			min = series.Candles[i].Date.UnixMilli()
			t = series.Candles[i].Date
		}
	}
	return t
}

func (series *CandleSeries) SortDateAsc() {
	sort.Slice(series.Candles, func(i, j int) bool {
		return series.Candles[i].Date.Unix() < series.Candles[j].Date.Unix()
	})
}

func (series *CandleSeries) SortDateDesc() {
	sort.Slice(series.Candles, func(i, j int) bool {
		return series.Candles[i].Date.Unix() > series.Candles[j].Date.Unix()
	})
}

func (series *CandleSeries) UpdateLimit() {
	if series.MaxCount > 0 && series.Len() > series.MaxCount {
		series.Candles = series.Candles[series.Len()-series.MaxCount : series.Len()]
	}
}

//func (series *CandleSeries) ToPeriods(fromPeriod period.Period, toPeriod period.Period, limit int) (*CandleSeries, error) {
//	if toPeriod < fromPeriod {
//		return nil, errors.New("sadece üst zaman dilimine çevrilebilir")
//	}
//	if toPeriod == fromPeriod {
//		return *&series, nil
//	}
//
//	periodSeries := &CandleSeries{
//		MaxCount: limit,
//	}
//	var curOpenDate *time.Time = nil
//	var curCloseDate *time.Time = nil
//	var curLastCandleDate *time.Time = nil
//	var _close *float64 = nil
//	var _open *float64 = nil
//	var _high *float64 = nil
//	var _low *float64 = nil
//	var _volume float64
//
//	for i := 0; i < series.Len(); i++ {
//		candle := series.Candles[i]
//		// Gap kontrolü için
//		if curOpenDate == nil {
//			// Yeni mum aç
//			curOpenDate = toPeriod.GetOpenDate(candle.Date)
//			curCloseDate = toPeriod.GetCloseDate(candle.Date)
//			curLastCandleDate = toPeriod.GetCloseDateForPeriod(candle.Date, fromPeriod)
//			_open = &candle.Open
//			_high = &candle.High
//			_low = &candle.Low
//			_volume = 0
//		}
//
//		_close = &candle.Close
//		if candle.High > *_high {
//			_high = &candle.High
//		}
//		if candle.Low < *_low {
//			_low = &candle.Low
//		}
//		_volume += candle.Volume
//
//		// Açık mumu kapat (Gap kontrollü)
//		if curLastCandleDate.UnixMilli() == candle.Date.UnixMilli() || (curCloseDate != nil && candle.Date.UnixMilli() > curCloseDate.UnixMilli()) {
//			periodSeries.AddCandle(&Candle{
//				Date:   curOpenDate,
//				Open:   *_open,
//				High:   *_high,
//				Low:    *_low,
//				Close:  *_close,
//				Volume: _volume,
//			})
//			curOpenDate = nil
//		}
//	}
//
//	return periodSeries, nil
//}

func (series *CandleSeries) Required(count int) bool {
	return series.Len() >= count
}

func (series *CandleSeries) ToPeriods(fromPeriod period.Period, toPeriod period.Period, limit int) (*CandleSeries, error) {
	if toPeriod < fromPeriod {
		return nil, errors.New("sadece üst zaman dilimine çevrilebilir")
	}
	if toPeriod == fromPeriod {
		clone := *&series
		clone.MaxCount = limit
		return clone, nil
	}

	periodSeries := &CandleSeries{
		MaxCount: limit,
	}
	var fillCandle *Candle = nil
	var lastSubCandleDate int64 = 0

	for i := 0; i < series.Len(); i++ {
		candle := series.Candles[i]
		// Gap kontrolü için
		if fillCandle == nil || toPeriod.IsTriggerTime(candle.Date) || candle.Date.UnixMilli() > lastSubCandleDate {
			// Yeni mum aç
			fillCandle = &Candle{
				Date:   *&candle.Date,
				Open:   candle.Open,
				High:   candle.High,
				Low:    candle.Low,
				Close:  candle.Close,
				Volume: candle.Volume,
			}
			periodSeries.AddCandle(fillCandle)
			lastSubCandleDate = toPeriod.GetCloseDateForPeriod(candle.Date, fromPeriod).UnixMilli()
			continue
		}

		if candle.High > fillCandle.High {
			fillCandle.High = candle.High
		}
		if candle.Low < fillCandle.Low {
			fillCandle.Low = candle.Low
		}
		fillCandle.Volume += candle.Volume
		fillCandle.Close = candle.Close
	}

	return periodSeries, nil
}

func (series *CandleSeries) ToPeriodLastCandle(toPeriod period.Period) *Candle {
	if series.Len() == 0 {
		return nil
	}
	// Tersine döngü - Get
	toCandle := &Candle{}
	for i := 0; i < series.Len(); i++ {
		c := series.Get(i)
		if i == 0 {
			toCandle.Close = c.Close
			toCandle.Low = c.Low
			toCandle.Date = toPeriod.GetOpenDate(c.Date)
		} else {
			if toPeriod.GetOpenDate(c.Date).UnixMilli() != toCandle.Date.UnixMilli() {
				break
			}
		}
		if c.High > toCandle.High {
			toCandle.High = c.High
		}
		if c.Low < toCandle.Low {
			toCandle.Low = c.Low
		}
		toCandle.Open = c.Open // ters döngü
		toCandle.Volume += c.Volume
	}

	return toCandle
}

// TakeToIndexArr arr : Float64 çıktısı,
// count : seriesten fazlası istenildiğinde seriesten toplan alınan count/*
func (series *CandleSeries) TakeRevIndexArr(takeCount int, revIndex int, sourceType SourceType) (arr *CandleArr, count int) {
	seriesLen := len(series.Candles) - revIndex
	startIndex := seriesLen + 1 - (takeCount + revIndex)
	if startIndex < 0 {
		ret := make(CandleArr, 0)
		return &ret, 0
	}
	ret := make(CandleArr, seriesLen-startIndex)
	for i := startIndex; i < seriesLen; i++ {
		ret[i-startIndex] = series.Candles[i].GetSourcePtr(sourceType)
	}
	return &ret, len(ret)
}
