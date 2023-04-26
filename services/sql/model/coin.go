package model

type Coin struct {
	tableName struct{} `pg:"coins"`
	ID        int      `pg:"id,pk"`
	ShortName string   `pg:"short_name"`
	Name      string   `pg:"name"`
	Type      int      `pg:"type"`
}
