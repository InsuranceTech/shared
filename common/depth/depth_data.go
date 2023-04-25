//go:generate msgp -o=DepthData_Msgpack.go -tests=false
package depth

type DepthData struct {
	Symbol string        `msg:"Symbol"`
	Bids   []*DepthPrice `msg:"Bids"`
	Asks   []*DepthPrice `msg:"Asks"`
}
