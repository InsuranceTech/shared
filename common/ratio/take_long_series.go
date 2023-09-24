//go:generate msgp -o=take_long_series_Msgpack.go -tests=false
package ratio

import "github.com/InsuranceTech/shared/common/symbol"

type TakeLongDataSeries struct {
	Symbol *symbol.Symbol
	Series []*TakeLongData
}
