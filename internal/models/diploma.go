package models

type Diploma struct {
	StudentID   int    `json:"student_id"`
	Theme       string `json:"theme"`
	Description string `json:"description"`
}
