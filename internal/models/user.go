package models

type User struct {
	ID           int    `json:"student_id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Birthday     string `json:"bitrhday"`
	Email        string `json:"email"`
	Password     string
	PasswordHash string
}

type Student struct {
	User
	GroupID int `json:"group_id"`
}

type Mentor struct {
	User
	Groups []Group
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
