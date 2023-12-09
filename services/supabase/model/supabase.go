package model

import "time"

type SupaBaseModel struct {
	Calculations int8      `json:"calculations"`
	Time         time.Time `json:"time"`
}
