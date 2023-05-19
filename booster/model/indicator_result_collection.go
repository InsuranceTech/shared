//go:generate msgp -o=indicator_result_collection_msgpack.go -tests=false
package model

import (
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/log"
	"github.com/InsuranceTech/shared/scanner"
	"sync"
)

type IndicatorResultCollection struct {
	Results map[period.Period][]*IndicatorResult

	// Gelişmiş filtreleme için
	// Period > FuncName > ExchangeType > Symbol[] indexleme
	pfs      map[period.Period]map[string]map[symbol.ExchangeType][]*IndicatorResult
	LastTime int64
	mutex    sync.Mutex
}

func CreateIndicatorResultCollection() *IndicatorResultCollection {
	i := &IndicatorResultCollection{
		mutex:   sync.Mutex{},
		Results: make(map[period.Period][]*IndicatorResult),
	}
	return i
}

func (c *IndicatorResultCollection) Append(result *IndicatorResult) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	if c.Results == nil {
		c.Results = make(map[period.Period][]*IndicatorResult)
	}
	_, ok := c.Results[result.Symbol.Period]
	if !ok {
		c.Results[result.Symbol.Period] = make([]*IndicatorResult, 0)
	}
	c.Results[result.Symbol.Period] = append(c.Results[result.Symbol.Period], result)
}

func (c *IndicatorResultCollection) Indexes() {
	// Period > FuncName > ExchangeType > Symbol[] indexleme
	c.pfs = make(map[period.Period]map[string]map[symbol.ExchangeType][]*IndicatorResult, 0)

	for keyPeriod, _ := range c.Results {
		for _, r := range c.Results[keyPeriod] {
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
}

func (c *IndicatorResultCollection) GetIndicatorsPFE(period period.Period, funcName string, exchangeType symbol.ExchangeType) []*IndicatorResult {
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

func (c *IndicatorResultCollection) GetIndicatorsPE(period period.Period, funcName string, exchangeType symbol.ExchangeType) []*IndicatorResult {
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

func (c *IndicatorResultCollection) GetIndicatorsP(period period.Period) []*IndicatorResult {
	// Period
	_, ok := c.Results[period]
	if ok == true {
		return c.Results[period]
	}
	return make([]*IndicatorResult, 0)
}

func (c *IndicatorResultCollection) Filters(exchangeType symbol.ExchangeType, period period.Period, conditions []*scanner.Condition) []*IndicatorResult {
	result := make([]*IndicatorResult, 0)

	if len(conditions) == 0 {
		return result
	}

	// Koşullar "VE" bağlacıyla bağlandığı için ilk koşul elemanını tara ve firsResults'a ekle
	firstCondition := conditions[0]
	firsResults := make([]*IndicatorResult, 0)
	items := c.GetIndicatorsPFE(period, firstCondition.FuncName, exchangeType)
	for _, item := range items {
		ok, err := item.CheckCondition(firstCondition)
		if err != nil {
			log.Error("Condition error", err)
		}
		if ok {
			firsResults = append(firsResults, item)
		}
	}

	// Diğer koşul elamanlarıda koşullara uyuyorsa result'a ekle
	if len(conditions) == 1 {
		result = firsResults
	} else {
		for _, item := range firsResults {
			allOk := true
			for i := 1; i < len(conditions); i++ {
				condition := conditions[i]
				ok, err := item.CheckCondition(condition)
				if err != nil {
					log.Error("Condition error", err)
				}
				if !ok {
					allOk = false
					break
				}
			}
			if allOk {
				result = append(result, item)
			}
		}
	}

	return result
}
