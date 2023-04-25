package common

import (
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
)

type PeriodicCandleSeries struct {
	Symbol        *symbol.Symbol
	PeriodCandles map[period.Period]*CandleSeries // Test ortamında JSON çıktısı alabilmek için public
}

func CreatePeriodicCandleSeries(symbol symbol.Symbol) *PeriodicCandleSeries {
	symbol.Period = period.NonePeriod
	return &PeriodicCandleSeries{
		Symbol:        &symbol,
		PeriodCandles: make(map[period.Period]*CandleSeries, 0),
	}
}

func (p *PeriodicCandleSeries) Get(period period.Period) *CandleSeries {
	return p.PeriodCandles[period]
}

func (p *PeriodicCandleSeries) Has(period period.Period) bool {
	_, ok := p.PeriodCandles[period]
	return ok
}

func (p *PeriodicCandleSeries) AddCandle(period period.Period, candle *Candle) bool {
	series := p.Get(period)
	if series == nil {
		return false
	}
	series.AddCandle(candle)
	return true
}

func (p *PeriodicCandleSeries) InsertCandle(period period.Period, candle *Candle) bool {
	series := p.Get(period)
	if series == nil {
		return false
	}
	series.InsertCandle(candle)
	return true
}

func (p *PeriodicCandleSeries) CreatePeriodSource(period period.Period, source *CandleSeries) bool {
	if p.Has(period) {
		return false
	}
	p.PeriodCandles[period] = source
	return true
}

func (p *PeriodicCandleSeries) CreatePeriod(period period.Period, limit int, candles []*Candle) bool {
	if p.Has(period) {
		return false
	}
	if candles == nil {
		candles = make([]*Candle, 0)
	}
	p.PeriodCandles[period] = &CandleSeries{
		MaxCount: limit,
		Candles:  candles,
	}

	return true
}
