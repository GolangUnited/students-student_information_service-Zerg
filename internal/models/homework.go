package models

import "time"

type Homework struct {
	HomeworkID int
	Text       string
	Deadline   time.Time
}

type HomeworkGrade struct {
	UserID     int
	HomeworkID int
	Grade      int
}
