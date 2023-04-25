//go:generate msgp -o=DepthPrice_Msgpack.go -tests=false
package depth

type DepthPrice struct {
	// Fiyat
	Price float64
	// Alım veya satımlar
	Sale float64
}
