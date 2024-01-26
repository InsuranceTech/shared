package model

import (
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
)

type Symbol struct {
	tableName      struct{} `pg:"symbols"`
	ID             int      `pg:"id,pk"`
	Name           string   `pg:"name"`
	BaseID         int      `pg:"base_id"`
	BaseCoin       *Coin    `pg:"rel:has-one,fk:base_id"`
	QuoteID        int      `pg:"quote_id"`
	QuoteCoin      *Coin    `pg:"rel:has-one,fk:quoteF_id"`
	ExchangeType   int      `pg:"exchange_type"`
	Status         string   `pg:"status"`
	Enabled        bool     `pg:"enabled,use_zero"`
	MinLot         float64  `pg:"min_lot"`
	StepSize       float64  `pg:"step_size"`
	MaxLot         float64  `pg:"max_lot"`
	MinPrice       float64  `pg:"min_price"`
	TickSize       float64  `pg:"tick_size"`
	MaxPrice       float64  `pg:"max_price"`
	MinNotional    float64  `pg:"min_notional"`
	CanBooster     bool     `pg:"can_booster,use_zero"`
	CanWhaleHunter bool     `pg:"can_whalehunter,use_zero"`
	CanScanner     bool     `pg:"can_scanner,use_zero"`
	CanTicker      bool     `pg:"can_ticker,use_zero"`
}

func (s *Symbol) ToSymbol() *symbol.Symbol {
	return symbol.NewSymbol(symbol.ExchangeType(s.ExchangeType), s.Name, period.NonePeriod)
}

func (s *Symbol) ToSymbolPeriod(period period.Period) *symbol.Symbol {
	return symbol.NewSymbol(symbol.ExchangeType(s.ExchangeType), s.Name, period)
}
