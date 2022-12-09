package models

type Contact struct {
	ID            int    `json:"contact_id"`
	UserID        int    `json:"user_id"`
	ContactTypeID int    `json:"contact_type_id"`
	ContactValue  string `json:"contact_value"`
}

type GroupContact struct {
	ID            int    `json:"group_contact_id"`
	GroupID       int    `json:"group_id"`
	ContactTypeID int    `json:"contact_type_id"`
	ContactValue  string `json:"contact_value"`
}
