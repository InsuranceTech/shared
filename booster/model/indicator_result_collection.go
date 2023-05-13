//go:generate msgp -o=indicator_result_collection_msgpack.go -tests=false
package model

import (
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
)

type IndicatorResultCollection struct {
	Results []*IndicatorResult
	// Gelişmiş filtreleme için
	// Period > FuncName > ExchangeType > Symbol[] indexleme
	pfs      map[period.Period]map[string]map[symbol.ExchangeType][]*IndicatorResult
	LastTime int64
}

func (c *IndicatorResultCollection) Append(result *IndicatorResult) {
	c.Results = append(c.Results, result)
}

func (c *IndicatorResultCollection) Indexes() {
	// Period > FuncName > ExchangeType > Symbol[] indexleme
	c.pfs = make(map[period.Period]map[string]map[symbol.ExchangeType][]*IndicatorResult, 0)

	for _, r := range c.Results {
		// Period
		_, ok := c.pfs[r.Symbol.Period]
		if ok == false {
			c.pfs[r.Symbol.Period] = make(map[string]map[symbol.ExchangeType][]*IndicatorResult, 0)
		}
		// Period > FuncName
		_, ok = c.pfs[r.Symbol.Period][r.FuncName]
		if ok == false {
			c.pfs[r.Symbol.Period][r.FuncName] = make(map[symbol.ExchangeType][]*IndicatorResult, 0)
		}
		// Period > FuncName > ExchangeType
		_, ok = c.pfs[r.Symbol.Period][r.FuncName][r.Symbol.Exchange]
		if ok == false {
			c.pfs[r.Symbol.Period][r.FuncName][r.Symbol.Exchange] = make([]*IndicatorResult, 0)
		}

		c.pfs[r.Symbol.Period][r.FuncName][r.Symbol.Exchange] = append(c.pfs[r.Symbol.Period][r.FuncName][r.Symbol.Exchange], r)
	}
}

func (c *IndicatorResultCollection) GetIndicators(period period.Period, funcName string, exchangeType symbol.ExchangeType) []*IndicatorResult {
	// Period
	_, ok := c.pfs[period]
	if ok == true {
		// Period > FuncName
		_, ok = c.pfs[period][funcName]
		if ok == true {
			// Period > FuncName > ExchangeType
			_, ok = c.pfs[period][funcName][exchangeType]
			if ok == true {
				return c.pfs[period][funcName][exchangeType]
			}
		}
	}
	return make([]*IndicatorResult, 0)
}
