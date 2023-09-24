//go:generate msgp -o=take_long_Msgpack.go -tests=false
package ratio

type TakeLongData struct {
	BuySellRatio string `json:"buySellRatio" msg:"buySellRatio"`
	SellVol      string `json:"sellVol" msg:"SellVol"`
	BuyVol       string `json:"buyVol" msg:"BuyVol"`
	Timestamp    int64  `json:"timestamp" msg:"Timestamp"`
}
