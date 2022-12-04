package models

type Group struct {
	GroupID  int    `json:"group_id"`
	CourseID int    `json:"course_id"`
	MentorID int    `json:"mentor_id"`
	Title    string `json:"title"`
}
