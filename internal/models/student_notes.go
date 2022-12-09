package models

type StudentNotes struct {
	StudentNoteID     int    `json:"student_note_id"`
	StudentID         int    `json:"student_id"`
	StudentNoteTypeID int    `json:"student_note_type_id"`
	Note              string `json:"student_note"`
}
