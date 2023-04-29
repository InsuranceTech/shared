//go:generate msgp -o=indicator_result_msgpack.go -tests=false
package model

import (
	"github.com/InsuranceTech/shared/common/symbol"
	"time"
)

type IndicatorResult struct {
	Symbol      *symbol.Symbol
	IndicatorID int64
	FuncName    string
	Values      map[string][]float64
	CandleTime  *time.Time
	UpdateTime  *time.Time
	Signal      int16
}
