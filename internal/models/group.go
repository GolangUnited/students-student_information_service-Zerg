package models

type GroupData struct {
	GroupID  int            `json:"group_id"`
	CourseID int            `json:"course_id"`
	MentorID int            `json:"mentor_id"`
	Title    string         `json:"title"`
	Contacts []GroupContact `json:"contacts"`
}

type Group struct {
	GroupData
	PrimaryContacatID int `json:"primary_contact_id"`
}
