package model

type UserLangCode struct {
	tableName struct{} `pg:"auth.users"`
	ID        int64    `pg:"id,pk" json:"id,omitempty"`
	UserLang  string   `pg:"lang_code" json:"lang_code,omitempty"`
}
