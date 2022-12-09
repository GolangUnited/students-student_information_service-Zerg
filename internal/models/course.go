package models

import "time"

type Course struct {
	CourseID       int       `json:"course_id"`
	Title          string    `json:"title"`
	StartDate      time.Time `json:"start_date"`
	EndDate        time.Time `json:"end_date"`
	CourseStatusID int       `json:"course_status_id"`
}

type CourseStatus struct {
	CourseStatusID int    `json:"course_status_id"`
	CourseStatus   string `json:"course_status"`
}

type CourseInterview struct {
	CourseInterviewID int `json:"course_interview_id"`
	CourseID          int `json:"course_id"`
	InterviewCount    int `json:"interview_count"`
	AverageScore      int `json:"average_score"`
	MaxScore          int `json:"max_score"`
	MinScore          int `json:"min_score"`
}
