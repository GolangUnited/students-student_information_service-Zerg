package models

type User struct {
	ID           int
	FirstName    string
	LastName     string
	Birthday     string
	Email        string
	PasswordHash string
}

type Student struct {
	User
	GroupID int
}
