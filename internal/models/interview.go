package models

import (
	"time"
)

type Interview struct {
	StudentID int
	MentorID  int
	Date      time.Time
	Mark      int
	Notes     string
}
