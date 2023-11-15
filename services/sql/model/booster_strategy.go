package model

type BoosterStrategy struct {
	tableName           struct{} `pg:"dev.booster_strategy"`
	ID                  int64    `pg:"id,pk" json:"id,omitempty"`
	UserID              int      `pg:"user_id" json:"user_id,omitempty"`
	Title               string   `pg:"title" json:"title,omitempty"`
	Indicators          []int64  `pg:"indicators,array" json:"indicators,omitempty"`
	IndicatorsFuncNames []string `pg:"indicator_func_names,array" json:"indicator_func_names,omitempty"`
}
