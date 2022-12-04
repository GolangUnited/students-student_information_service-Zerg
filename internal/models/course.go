package models

import "time"

type Course struct {
	CourseID  int       `json:"course_id"`
	Title     string    `json:"title"`
	StartDate time.Time `json:"start_date"`
	EndDate   time.Time `json:"end_date"`
}
