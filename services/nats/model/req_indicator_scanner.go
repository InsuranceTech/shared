package model

import (
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/scanner"
)

type ReqIndicatorScanner struct {
	ExchangeType symbol.ExchangeType  `json:"exchange_type"`
	Period       period.Period        `json:"period"`
	Conditions   []*scanner.Condition `json:"conditions"`
}
