package model

type UserLangCode struct {
	tableName struct{} `pg:"auth.users"`
	UserLang  string   `pg:"lang_code" json:"lang_code,omitempty"`
}
