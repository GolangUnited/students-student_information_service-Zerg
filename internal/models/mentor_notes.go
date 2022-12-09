package models

type MentorNotes struct {
	MentorNoteID int    `json:"mentor_note_id"`
	MentorID     int    `json:"mentor_id"`
	StudentID    int    `json:"student_id"`
	Note         string `json:"mentor_note"`
}
