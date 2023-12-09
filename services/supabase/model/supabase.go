package model

import "time"

type SupaBaseModel struct {
	Calculations string    `json:"calculations"`
	Time         time.Time `json:"time"`
}
