package model

type Indicator struct {
	tableName      struct{} `pg:"indicators"`
	ID             int      `pg:"id,pk"`
	Name           string   `pg:"name"`
	Description    string   `pg:"description"`
	FuncName       string   `pg:"func_name"`
	Status         bool     `pg:"status"`
	ExcludeSymbols []int    `pg:"exclude_symbols"`
}
