package models

import "time"

type Homework struct {
	HomeworkID int       `json:"homework_id"`
	Text       string    `json:"homework_text"`
	Deadline   time.Time `json:"deadline"`
}

type HomeworkGrade struct {
	HomeworkGradeID int `json:"homework_grade_id"`
	StudentID       int `json:"student_id"`
	HomeworkID      int `json:"homework_id"`
	Grade           int `json:"grade"`
}
