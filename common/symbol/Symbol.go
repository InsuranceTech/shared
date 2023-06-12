//msgp:tuple Symbol
//go:generate msgp -o=Symbol_Msgpack.go -tests=false
package symbol

import (
	"github.com/InsuranceTech/shared/common/period"
	"strconv"
	"strings"
)

type ExchangeType int

const (
	BINANCE_SPOT        ExchangeType = 1
	BINANCE_USDT_FUTURE ExchangeType = 2
	BINANCE_COIN_FUTURE ExchangeType = 3
)

type Symbol struct {
	SymbolName string        `json:"symbolName"`
	Exchange   ExchangeType  `json:"exchange"`
	Period     period.Period `json:"period"`
}

// ToString : Çıktı: EXCHANGE:SYMBOL:PERIOD (StringBuilder sebebi: https://freshman.tech/snippets/go/string-concatenation/ )
func (s *Symbol) ToString() string {
	sb := strings.Builder{}
	sb.WriteString(s.ToStringNoPeriod())
	if s.Period != period.NonePeriod {
		sb.WriteString(":")
		sb.WriteString(s.Period.MinuteStr())
	}
	return sb.String()
}

// ToStringNoPeriod : Çıktı: EXCHANGE:SYMBOL (StringBuilder sebebi: https://freshman.tech/snippets/go/string-concatenation/ )
func (s *Symbol) ToStringNoPeriod() string {
	sb := strings.Builder{}
	sb.WriteString(s.Exchange.ToString())
	sb.WriteString(":")
	sb.WriteString(s.SymbolName)
	return sb.String()
}

func (ex *ExchangeType) ToString() string {
	if *ex == BINANCE_SPOT {
		return "BINANCE_SPOT"
	} else if *ex == BINANCE_USDT_FUTURE {
		return "BINANCE_USDT_FUTURE"
	} else if *ex == BINANCE_COIN_FUTURE {
		return "BINANCE_COIN_FUTURE"
	} else {
		return "UNK"
	}
}

func NewSymbol(exchangeType ExchangeType, symbol string, period period.Period) *Symbol {
	return &Symbol{
		SymbolName: symbol,
		Period:     period,
		Exchange:   exchangeType,
	}
}

func NewBinanceSpotSymbol(symbol string, period period.Period) *Symbol {
	return &Symbol{
		SymbolName: symbol,
		Period:     period,
		Exchange:   BINANCE_SPOT,
	}
}

func NewBinanceUsdtFutureSymbol(symbol string, period period.Period) *Symbol {
	return &Symbol{
		SymbolName: symbol,
		Period:     period,
		Exchange:   BINANCE_USDT_FUTURE,
	}
}

func NewBinanceCoinFutureSymbol(symbol string, period period.Period) *Symbol {
	return &Symbol{
		SymbolName: symbol,
		Period:     period,
		Exchange:   BINANCE_COIN_FUTURE,
	}
}

/*
ParseSymbolEx : BINANCE_SPOT:BTCUSDT:5 yapısını symbole çevir
EXCHANGE:SYMBOL
EXCHANGE:SYMBOL:PERIOD
*/
func ParseSymbolEx(symbolExName string) (symbol *Symbol, ok bool) {
	splits := strings.Split(symbolExName, ":")

	if len(splits) == 2 {
		// EXCHANGE:SYMBOL
		exchange := stringToExchangeType(splits[0])
		symbolName := splits[1]
		return &Symbol{
			SymbolName: symbolName,
			Exchange:   exchange,
			Period:     period.NonePeriod,
		}, true
	} else if len(splits) == 3 {
		// EXCHANGE:SYMBOL:PERIOD
		exchange := stringToExchangeType(splits[0])
		symbolName := splits[1]
		p := stringToPeriod(splits[2])
		return &Symbol{
			SymbolName: symbolName,
			Exchange:   exchange,
			Period:     p,
		}, true
	} else {
		// Unknown
		return &Symbol{
			SymbolName: "",
			Exchange:   0,
			Period:     period.NonePeriod,
		}, false
	}
}

func stringToExchangeType(exchangeName string) ExchangeType {
	exchangeName = strings.ToUpper(exchangeName)
	if exchangeName == "BINANCE_SPOT" {
		return BINANCE_SPOT
	}
	if exchangeName == "BINANCE_USDT_FUTURE" {
		return BINANCE_USDT_FUTURE
	}
	if exchangeName == "BINANCE_COIN_FUTURE" {
		return BINANCE_COIN_FUTURE
	}
	return 0
}

func stringToPeriod(periodStr string) period.Period {
	periodInt, err := strconv.Atoi(periodStr)
	if err != nil {
		return period.NonePeriod
	}
	return period.Period(periodInt)
}

func (s *Symbol) Clone() *Symbol {
	clone := *s
	return &clone
}

func (s *Symbol) CloneP(p period.Period) *Symbol {
	clone := s.Clone()
	clone.Period = p
	return clone
}
