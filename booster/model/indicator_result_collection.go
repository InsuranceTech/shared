//go:generate msgp -o=indicator_result_collection_msgpack.go -tests=false
package model

import (
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/log"
	"github.com/InsuranceTech/shared/scanner"
	"github.com/spf13/cast"
	"sync"
)

type IndicatorResultCollection struct {
	Results map[period.Period][]*IndicatorResult

	// Gelişmiş filtreleme için
	// Period > FuncName > ExchangeType > Symbol[] indexleme
	pfe map[period.Period]map[string]map[symbol.ExchangeType][]*IndicatorResult
	// Period > FuncName > ExchangeType > SymbolName indexleme
	pfes     map[period.Period]map[string]map[symbol.ExchangeType]map[string]*IndicatorResult
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
	c.pfe = make(map[period.Period]map[string]map[symbol.ExchangeType][]*IndicatorResult, 0)
	// Period > FuncName > ExchangeType > SymbolName indexleme
	c.pfes = make(map[period.Period]map[string]map[symbol.ExchangeType]map[string]*IndicatorResult, 0)

	for keyPeriod, _ := range c.Results {
		for _, r := range c.Results[keyPeriod] {
			// pfe
			// Period
			_, ok := c.pfe[r.Symbol.Period]
			if ok == false {
				c.pfe[r.Symbol.Period] = make(map[string]map[symbol.ExchangeType][]*IndicatorResult, 0)
			}
			// Period > FuncName
			_, ok = c.pfe[r.Symbol.Period][r.FuncName]
			if ok == false {
				c.pfe[r.Symbol.Period][r.FuncName] = make(map[symbol.ExchangeType][]*IndicatorResult, 0)
			}
			// Period > FuncName > ExchangeType
			_, ok = c.pfe[r.Symbol.Period][r.FuncName][r.Symbol.Exchange]
			if ok == false {
				c.pfe[r.Symbol.Period][r.FuncName][r.Symbol.Exchange] = make([]*IndicatorResult, 0)
			}

			c.pfe[r.Symbol.Period][r.FuncName][r.Symbol.Exchange] = append(c.pfe[r.Symbol.Period][r.FuncName][r.Symbol.Exchange], r)

			// -------------------
			// pfes
			// Period
			_, ok = c.pfes[r.Symbol.Period]
			if ok == false {
				c.pfes[r.Symbol.Period] = make(map[string]map[symbol.ExchangeType]map[string]*IndicatorResult, 0)
			}
			// Period > FuncName
			_, ok = c.pfes[r.Symbol.Period][r.FuncName]
			if ok == false {
				c.pfes[r.Symbol.Period][r.FuncName] = make(map[symbol.ExchangeType]map[string]*IndicatorResult, 0)
			}
			// Period > FuncName > ExchangeType
			_, ok = c.pfes[r.Symbol.Period][r.FuncName][r.Symbol.Exchange]
			if ok == false {
				c.pfes[r.Symbol.Period][r.FuncName][r.Symbol.Exchange] = make(map[string]*IndicatorResult, 0)
			}
			// Period > FuncName > ExchangeType > SymbolName
			c.pfes[r.Symbol.Period][r.FuncName][r.Symbol.Exchange][r.Symbol.SymbolName] = r
		}

	}
}

func (c *IndicatorResultCollection) GetIndicatorsPFES(funcName string, symbol *symbol.Symbol) *IndicatorResult {
	// Period
	_, ok := c.pfe[symbol.Period]
	if ok == true {
		// Period > FuncName
		_, ok = c.pfe[symbol.Period][funcName]
		if ok == true {
			// Period > FuncName > ExchangeType
			_, ok = c.pfes[symbol.Period][funcName][symbol.Exchange]
			if ok == true {
				// Period > FuncName > ExchangeType > SymbolName
				_, ok = c.pfes[symbol.Period][funcName][symbol.Exchange][symbol.SymbolName]
				if ok == true {
					return c.pfes[symbol.Period][funcName][symbol.Exchange][symbol.SymbolName]
				}
			}
		}
	}
	return nil
}

func (c *IndicatorResultCollection) GetIndicatorsPFE(period period.Period, funcName string, exchangeType symbol.ExchangeType) []*IndicatorResult {
	// Period
	_, ok := c.pfe[period]
	if ok == true {
		// Period > FuncName
		_, ok = c.pfe[period][funcName]
		if ok == true {
			// Period > FuncName > ExchangeType
			_, ok = c.pfe[period][funcName][exchangeType]
			if ok == true {
				return c.pfe[period][funcName][exchangeType]
			}
		}
	}
	return make([]*IndicatorResult, 0)
}

func (c *IndicatorResultCollection) GetIndicatorsPE(period period.Period, funcName string, exchangeType symbol.ExchangeType) []*IndicatorResult {
	// Period
	_, ok := c.pfe[period]
	if ok == true {
		// Period > FuncName
		_, ok = c.pfe[period][funcName]
		if ok == true {
			// Period > FuncName > ExchangeType
			_, ok = c.pfe[period][funcName][exchangeType]
			if ok == true {
				return c.pfe[period][funcName][exchangeType]
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

func toResultItem(condition *scanner.Condition, result *IndicatorResult) *scanner.ResultItem {
	i := &scanner.ResultItem{
		Symbol: result.Symbol,
		Outs:   make([]*scanner.ResulIndicatortItem, 0),
	}
	appendOutToResultItem(condition, result, i)
	return i
}

func appendOutToResultItem(condition *scanner.Condition, result *IndicatorResult, item *scanner.ResultItem) {
	val := ""
	outName := ""

	if condition.IsSourceFixedValue() {
		v, _ := result.GetLastValue(condition.Target)
		val = cast.ToString(v)
		outName = condition.Target
	} else {
		v, _ := result.GetLastValue(condition.Source)
		val = cast.ToString(v)
		outName = condition.Source
	}
	item.Outs = append(item.Outs, &scanner.ResulIndicatortItem{
		IndicatorID: condition.IndicatorID,
		FuncName:    condition.FuncName,
		OutName:     outName,
		Value:       val,
	})
}

func (c *IndicatorResultCollection) Filters(exchangeType symbol.ExchangeType, period period.Period, conditions []*scanner.Condition) []*scanner.ResultItem {
	result := make([]*scanner.ResultItem, 0)
	firstResults := make([]*scanner.ResultItem, 0)

	if len(conditions) == 0 {
		return result
	}

	// Koşullar "VE" bağlacıyla bağlandığı için ilk koşul elemanını tara ve firstItems'a ekle
	firstCondition := conditions[0]
	firstItems := make([]*IndicatorResult, 0)
	items := c.GetIndicatorsPFE(period, firstCondition.FuncName, exchangeType)
	for _, item := range items {
		ok, err := item.CheckCondition(firstCondition)
		if err != nil {
			log.Error("Condition error", err)
		}
		if ok {
			firstItems = append(firstItems, item)
			firstResults = append(firstResults, toResultItem(firstCondition, item))
		}
	}

	// Diğer koşul elamanlarıda koşullara uyuyorsa result'a ekle
	if len(conditions) == 1 {
		return firstResults
	} else {
		for _fi := 0; _fi < len(firstItems); _fi++ {
			firstItemSymbol := firstItems[_fi].Symbol
			firstResult := firstResults[_fi]
			allOk := true
			for i := 1; i < len(conditions); i++ {
				condition := conditions[i]
				scanItem := c.GetIndicatorsPFES(condition.FuncName, firstItemSymbol)
				if scanItem == nil {
					// Paritenin bu gösterge için verisi yok ?
					allOk = false
					break
				}
				ok, err := scanItem.CheckCondition(condition)
				if err != nil {
					log.Error("Condition error", err)
				}
				if !ok {
					allOk = false
					break
				}
				appendOutToResultItem(condition, scanItem, firstResult)
			}
			if allOk {
				result = append(result, firstResult)
			}
		}
	}

	return result
}
