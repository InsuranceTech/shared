package model

type UserInfo struct {
	tableName struct{}    `pg:"auth.users"`
	ID        int64       `pg:"id,pk" json:"id,omitempty"`
	UserLang  string      `pg:"lang_code" json:"lang_code,omitempty"`
	FcmTokens []*FcmToken `pg:"rel:has-many,join_fk:user_id"`
}

type FcmToken struct {
	tableName struct{} `pg:"fcm_tokens"`
	ID        int64    `pg:"id,pk"`
	UserID    int64    `pg:"user_id"`
	TypeOf    string   `pg:"type_of"`
	Token     string   `pg:"fcm_token"`
	IsAct     bool     `pg:"is_act"`
}
