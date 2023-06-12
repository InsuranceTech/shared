package scanner

type ResulIndicatortItem struct {
	IndicatorID int    `json:"indicator_id"`
	FuncName    string `json:"func_name"`
	OutName     string `json:"out_name"`
	Value       string `json:"value"`
}
