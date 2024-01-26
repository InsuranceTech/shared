package model

type UserLangCode struct {
	tableName struct{} `pg:"auth.users"`
	ID        int32    `pg:"id,pk"`
	UserLang  string   `pg:"lang_code" json:"lang_code,omitempty"`
}
