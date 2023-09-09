package extension

// import cycle yüzünden extension methodlar burada
import (
	"github.com/InsuranceTech/shared/common/period"
	"github.com/InsuranceTech/shared/common/symbol"
	"github.com/InsuranceTech/shared/services/sql/model"
)

func FromModelSymbols(symbols []*model.Symbol) []*symbol.Symbol {
	retSymbols := make([]*symbol.Symbol, 0)
	if len(symbols) == 0 {
		return retSymbols
	}
	for _, s := range symbols {
		switch s.ExchangeType {
		case int(symbol.BINANCE_SPOT):
			retSymbols = append(retSymbols, symbol.NewBinanceSpotSymbol(s.Name, period.NonePeriod))
			break
		case int(symbol.BINANCE_USDT_FUTURE):
			retSymbols = append(retSymbols, symbol.NewBinanceUsdtFutureSymbol(s.Name, period.NonePeriod))
			break
		case int(symbol.BINANCE_COIN_FUTURE):
			retSymbols = append(retSymbols, symbol.NewBinanceCoinFutureSymbol(s.Name, period.NonePeriod))
			break
		}
	}
	return retSymbols
}
