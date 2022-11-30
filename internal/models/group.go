package models

import (
	"time"
)

type Group struct {
	ID        int       `json:"group_id"`
	Mentor    int       `json:"mentor_id"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
}
