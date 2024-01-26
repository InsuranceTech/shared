package model

type FcmToken struct {
	tableName struct{} `pg:"auth.fcm_tokens"`
	ID        int64    `pg:"id,pk" json:"id,omitempty"`
	UserID    int      `pg:"user_id" json:"user_id,omitempty"`
	TypeOf    string   `pg:"type_of" json:"type_of,omitempty"`
	Token     string   `pg:"fcm_token" json:"fcm_token,omitempty"`
	IsAct     bool     `pg:"is_act" json:"is_act,omitempty"`
}
