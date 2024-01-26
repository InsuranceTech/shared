package model

import (
	"github.com/InsuranceTech/shared/scanner"
	"time"
)

type AlarmIndicator struct {
	tableName       struct{}            `pg:"indicators"`
	ID              int                 `pg:"id,pk" json:"id,omitempty"`
	UserID          int                 `pg:"user_id" json:"user_id,omitempty"`
	Name            string              `pg:"name" json:"name,omitempty"`
	Symbol          string              `pg:"symbol" json:"symbol,omitempty"`
	TriggerCount    int                 `pg:"trigger_count" json:"trigger_count,omitempty"`
	LastTriggerTime time.Time           `pg:"last_trigger" json:"last_trigger_time"`
	Conditions      []scanner.Condition `pg:"conditions,array" json:"conditions,omitempty"`
	FcmTokens       []*FcmToken         `pg:"rel:has-many,fk:user_id"`
	UserLang        *UserLangCode       `pg:"rel:has-one,fk:user_id"`
}
