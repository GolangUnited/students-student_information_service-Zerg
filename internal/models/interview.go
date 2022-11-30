package models

import (
	"time"
)

type Interview struct {
	StudentID int       `json:"student_id"`
	MentorID  int       `json:"mentor_id"`
	Date      time.Time `json:"interview_date"`
	Mark      int       `json:"mark"`
	Notes     string    `json:"notes"`
}
