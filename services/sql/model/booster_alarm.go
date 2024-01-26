package model

import "time"

type EAlarmStatus int16

const (
	AS_Nothing EAlarmStatus = -99
	AS_Sell    EAlarmStatus = -1
	AS_Notr    EAlarmStatus = 0
	AS_Buy     EAlarmStatus = 1
)

type EAlarmChangeType int16

const (
	// AC_All Tüm renk değişimlerinde
	AC_All EAlarmChangeType = -99
	// AC_Sell Sadece kırmızıya döndüğünde
	AC_Sell EAlarmChangeType = -1
	// AC_Notr Sadece sarıya döndüğünde
	AC_Notr EAlarmChangeType = 0
	// AC_Buy Sadece yeşile döndüğünde
	AC_Buy EAlarmChangeType = 1
)

type BoosterAlarm struct {
	tableName       struct{}         `pg:"alarm.boosters"`
	ID              int              `pg:"id,pk" json:"id,omitempty"`
	UserID          int              `pg:"user_id" json:"user_id,omitempty"`
	Name            string           `pg:"name" json:"name,omitempty"`
	Symbol          string           `pg:"symbol" json:"symbol,omitempty"`
	ChangeType      EAlarmChangeType `pg:"change_type" json:"change_type,omitempty"` // -99 : Tüm değişim, -1: Kırmızı, 0: Sarı, 1: Yeşil
	TriggerCount    int              `pg:"trigger_count" json:"trigger_count,omitempty"`
	LastTriggerTime *time.Time       `pg:"last_trigger" json:"last_trigger_time"`
	StrategyID      int              `pg:"strategy_id" json:"strategy_id,omitempty"`
	NeedCount       int              `pg:"need_count" json:"need_count,omitempty"`
	Triggered       bool             `pg:"triggered,use_zero" json:"triggered,omitempty"`
	TriggeredStatus *EAlarmStatus    `pg:"triggered_status" json:"triggered_status,omitempty"`
	EndOf           *time.Time       `pg:"end_of" json:"end_of,omitempty"`
	Enable          bool             `pg:"enable,use_zero" json:"enable,omitempty"`
	BoosterStrategy *BoosterStrategy `pg:"rel:has-one,fk:strategy_id"`
	FcmTokens       []*FcmToken      `pg:"rel:has-many,join_fk:user_id"`
	UserInfo        *UserInfo        `pg:"rel:has-one,fk:user_id"`
}
