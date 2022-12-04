package models

type MentorNotes struct {
	MentorID  int    `json:"mentor_id"`
	StudentID int    `json:"student_id"`
	Note      string `json:"mentor_note"`
}
