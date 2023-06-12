package scanner

import "github.com/InsuranceTech/shared/common/symbol"

type ResultItem struct {
	Symbol *symbol.Symbol         `json:"symbol"`
	Outs   []*ResulIndicatortItem `json:"outs"`
}
