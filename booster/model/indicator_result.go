//go:generate msgp -o=indicator_result_msgpack.go -tests=false
package model

import (
	"errors"
	"github.com/InsuranceTech/shared/common"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/scanner"
	"math"
	"time"
)

type IndicatorResult struct {
	Symbol      *symbol.Symbol
	IndicatorID int64
	FuncName    string
	Values      map[string][]float64 // { OutName: [ SonVeri, ÖncekiVeri ] }
	LastCandle  *common.Candle
	PrevCandle  *common.Candle
	UpdateTime  *time.Time
	Signal      int16
}

func (i *IndicatorResult) CheckCondition(condition *scanner.Condition) (bool, error) {
	lastSourceValue, ok1 := GetLastSourceValue(condition, i)
	lastTargetValue, ok2 := GetLastSourceValue(condition, i)
	if !ok1 {
		return false, errors.New("cannot read Last Source Value")
	}
	if !ok2 {
		return false, errors.New("cannot read Last Target Value")
	}

	switch condition.Operator {
	case scanner.CO_Less:
		return lastSourceValue < lastTargetValue, nil
	case scanner.CO_LessOrEqual:
		return lastSourceValue <= lastTargetValue, nil
	case scanner.CO_Greater:
		return lastSourceValue > lastTargetValue, nil
	case scanner.CO_GreaterOrEqual:
		return lastSourceValue >= lastTargetValue, nil
	case scanner.CO_Range:
		return lastSourceValue >= lastTargetValue && lastSourceValue <= lastTargetValue, nil
	case scanner.CO_OutRange:
		return !(lastSourceValue >= lastTargetValue && lastSourceValue <= lastTargetValue), nil
	case scanner.CO_Equal:
		return equalFloat(lastSourceValue, lastTargetValue), nil
	case scanner.CO_NotEqual:
		return !equalFloat(lastSourceValue, lastTargetValue), nil
	case scanner.CO_CrossUp:
		prevSourceValue, ok1 := GetPrevSourceValue(condition, i)
		prevTargetValue, ok2 := GetPrevSourceValue(condition, i)
		if !ok1 {
			return false, errors.New("cannot read Prev Source Value")
		}
		if !ok2 {
			return false, errors.New("cannot read Prev Target Value")
		}
		return (lastSourceValue > lastTargetValue) && !(prevSourceValue > prevTargetValue), nil
	case scanner.CO_CrossDown:
		prevSourceValue, ok1 := GetPrevSourceValue(condition, i)
		prevTargetValue, ok2 := GetPrevSourceValue(condition, i)
		if !ok1 {
			return false, errors.New("cannot read Prev Source Value")
		}
		if !ok2 {
			return false, errors.New("cannot read Prev Target Value")
		}
		return (lastSourceValue < lastTargetValue) && !(prevSourceValue < prevTargetValue), nil
	default:
		return false, errors.New("unknown Operator Type")
	}
}

func (i *IndicatorResult) GetLastValue(outName string) (float64, bool) {
	if i.Values == nil {
		return 0, false
	}
	v, ok := i.Values[outName]
	if !ok || len(v) != 2 {
		return 0, false
	}
	return v[0], true
}

func (i *IndicatorResult) GetPrevValue(outName string) (float64, bool) {
	if i.Values == nil {
		return 0, false
	}
	v, ok := i.Values[outName]
	if !ok || len(v) != 2 {
		return 0, false
	}
	return v[1], true
}

func equalFloat(a, b float64) bool {
	tolerance := 0.1
	if diff := math.Abs(a - b); diff < tolerance {
		return true
	} else {
		return false
	}
}

func GetLastSourceValue(c *scanner.Condition, result *IndicatorResult) (float64, bool) {
	switch c.Source {
	case scanner.CO_FV_CANDLE_OPEN:
		if result.LastCandle != nil {
			return result.LastCandle.GetSource(common.ST_OPEN), true
		}
	case scanner.CO_FV_CANDLE_HIGH:
		if result.LastCandle != nil {
			return result.LastCandle.GetSource(common.ST_HIGH), true
		}
	case scanner.CO_FV_CANDLE_LOW:
		if result.LastCandle != nil {
			return result.LastCandle.GetSource(common.ST_LOW), true
		}
	case scanner.CO_FV_CANDLE_CLOSE:
		if result.LastCandle != nil {
			return result.LastCandle.GetSource(common.ST_CLOSE), true
		}
	default:
		return result.GetLastValue(c.Source)
	}
	return 0, false
}

func GetLastTargetValue(c *scanner.Condition, result *IndicatorResult) (float64, bool) {
	switch c.Target {
	case scanner.CO_FV_CANDLE_OPEN:
		if result.LastCandle != nil {
			return result.LastCandle.GetSource(common.ST_OPEN), true
		}
	case scanner.CO_FV_CANDLE_HIGH:
		if result.LastCandle != nil {
			return result.LastCandle.GetSource(common.ST_HIGH), true
		}
	case scanner.CO_FV_CANDLE_LOW:
		if result.LastCandle != nil {
			return result.LastCandle.GetSource(common.ST_LOW), true
		}
	case scanner.CO_FV_CANDLE_CLOSE:
		if result.LastCandle != nil {
			return result.LastCandle.GetSource(common.ST_CLOSE), true
		}
	default:
		return result.GetLastValue(c.Target)
	}
	return 0, false
}

func GetPrevSourceValue(c *scanner.Condition, result *IndicatorResult) (float64, bool) {
	switch c.Source {
	case scanner.CO_FV_CANDLE_OPEN:
		if result.PrevCandle != nil {
			return result.PrevCandle.GetSource(common.ST_OPEN), true
		}
	case scanner.CO_FV_CANDLE_HIGH:
		if result.PrevCandle != nil {
			return result.PrevCandle.GetSource(common.ST_HIGH), true
		}
	case scanner.CO_FV_CANDLE_LOW:
		if result.PrevCandle != nil {
			return result.PrevCandle.GetSource(common.ST_LOW), true
		}
	case scanner.CO_FV_CANDLE_CLOSE:
		if result.PrevCandle != nil {
			return result.PrevCandle.GetSource(common.ST_CLOSE), true
		}
	default:
		return result.GetPrevValue(c.Source)
	}
	return 0, false
}

func GetPrevTargetValue(c *scanner.Condition, result *IndicatorResult) (float64, bool) {
	switch c.Target {
	case scanner.CO_FV_CANDLE_OPEN:
		if result.PrevCandle != nil {
			return result.PrevCandle.GetSource(common.ST_OPEN), true
		}
	case scanner.CO_FV_CANDLE_HIGH:
		if result.PrevCandle != nil {
			return result.PrevCandle.GetSource(common.ST_HIGH), true
		}
	case scanner.CO_FV_CANDLE_LOW:
		if result.PrevCandle != nil {
			return result.PrevCandle.GetSource(common.ST_LOW), true
		}
	case scanner.CO_FV_CANDLE_CLOSE:
		if result.PrevCandle != nil {
			return result.PrevCandle.GetSource(common.ST_CLOSE), true
		}
	default:
		return result.GetPrevValue(c.Target)
	}
	return 0, false
}
