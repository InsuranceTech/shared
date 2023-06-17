package model

import (
	"github.com/InsuranceTech/shared/scanner"
	"time"
)

type AlarmIndicator struct {
	tableName       struct{}            `pg:"indicators"`
	ID              int                 `pg:"id,pk"`
	UserID          int                 `pg:"user_id"`
	Name            string              `pg:"name"`
	Symbol          string              `pg:"symbol"`
	TriggerCount    int                 `pg:"trigger_count"`
	LastTriggerTime time.Time           `pg:"last_trigger"`
	Conditions      []scanner.Condition `pg:"conditions,array"`
}
