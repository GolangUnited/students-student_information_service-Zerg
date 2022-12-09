package models

type Diploma struct {
	GroupID     int    `json:"group_id"`
	Theme       string `json:"theme"`
	Description string `json:"description"`
}
