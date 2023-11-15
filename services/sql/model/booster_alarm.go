package model

import "time"

type EAlarmStatus int

const (
	AS_Nothing EAlarmStatus = -99
	AS_Sell    EAlarmStatus = -1
	AS_Notr    EAlarmStatus = 0
	AS_Buy     EAlarmStatus = 1
)

type BoosterAlarm struct {
	tableName       struct{}         `pg:"alarm.boosters"`
	ID              int              `pg:"id,pk" json:"id,omitempty"`
	UserID          int              `pg:"user_id" json:"user_id,omitempty"`
	Name            string           `pg:"name" json:"name,omitempty"`
	Symbol          string           `pg:"symbol" json:"symbol,omitempty"`
	TriggerCount    int              `pg:"trigger_count" json:"trigger_count,omitempty"`
	LastTriggerTime *time.Time       `pg:"last_trigger" json:"last_trigger_time"`
	StrategyID      int              `pg:"strategy_id" json:"strategy_id,omitempty"`
	NeedCount       int              `pg:"need_count" json:"need_count,omitempty"`
	AlarmStatus     EAlarmStatus     `pg:"alarm_status" json:"alarm_status,omitempty"`
	EndOf           *time.Time       `pg:"end_of" json:"end_of,omitempty"`
	Enable          bool             `pg:"enable" json:"enable,omitempty"`
	BoosterStrategy *BoosterStrategy `pg:"rel:has-one,fk:strategy_id"`
}
