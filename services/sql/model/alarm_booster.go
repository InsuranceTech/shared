package model

import (
	"time"
)

type AlarmBooster struct {
	tableName       struct{}  `pg:"boosters"`
	ID              int       `pg:"id,pk" json:"id,omitempty"`
	UserID          int       `pg:"user_id" json:"user_id,omitempty"`
	Name            string    `pg:"name" json:"name,omitempty"`
	Symbol          string    `pg:"symbol" json:"symbol,omitempty"`
	TriggerCount    int       `pg:"trigger_count" json:"trigger_count,omitempty"`
	LastTriggerTime time.Time `pg:"last_trigger" json:"last_trigger_time"`
	WatchListID     int       `pg:"watch_list_id" json:"watch_list_id,omitempty"`
}
