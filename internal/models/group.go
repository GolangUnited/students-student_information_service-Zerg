package models

import (
	"time"
)

type Group struct {
	ID        int
	mentor    int
	dateStart time.Time
	dateEnd   time.Time
}
