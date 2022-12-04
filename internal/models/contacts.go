package models

type Contact struct {
	ID            int    `json:"contact_id"`
	PersonID      int    `json:"person_id"`
	ContactTypeID int    `json:"contact_type_id"`
	ContactValue  string `json:"contact_value"`
}

type GroupContact struct {
	ID            int    `json:"contact_id"`
	GroupID       int    `json:"group_id"`
	ContactTypeID int    `json:"contact_type_id"`
	ContactValue  string `json:"contact_value"`
}
