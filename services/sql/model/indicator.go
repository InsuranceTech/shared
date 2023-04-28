package model

type Indicator struct {
	tableName       struct{} `pg:"indicators"`
	ID              int64    `pg:"id,pk"`
	Category        []int64  `pg:"category,array"`
	Type            int16    `pg:"type_of"`
	Name            string   `pg:"name"`
	Description     string   `pg:"description"`
	FuncName        string   `pg:"func_name"`
	CanBooster      bool     `pg:"can_booster"`
	CanScanner      bool     `pg:"can_adv_filter"`
	CanAlarm        bool     `pg:"can_alarm"`
	CanBoosterTable bool     `pg:"can_booster_tbl"`
	Status          bool     `pg:"status"`
}
