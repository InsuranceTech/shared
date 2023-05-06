package model

type IndicatorOut struct {
	tableName   struct{} `pg:"indicator_outs"`
	ID          int32    `pg:"id,pk"`
	FuncName    string   `pg:"func_name"`
	IndicatorID int64    `pg:"indicator_id"`
	Name        string   `pg:"name"`
	Type        int32    `pg:"type"`
	ControlType int32    `pg:"control_type"`
	MinValue    float64  `pg:"min_value"`
	MaxValue    float64  `pg:"max_value"`
}
