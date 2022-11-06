package models

import (
	"time"
)

type Group struct {
	ID        int
	Mentor    int
	DateStart time.Time
	DateEnd   time.Time
}
