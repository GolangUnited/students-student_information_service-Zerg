package models

import (
	"time"
)

type Group struct {
	GroupID   int       `json:"group_id"`
	MentorID  int       `json:"mentor_id"`
	GroupName string    `json:"group_name"`
	DateStart time.Time `json:"date_start"`
	DateEnd   time.Time `json:"date_end"`
}
