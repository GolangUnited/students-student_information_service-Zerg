package models

type User struct {
	ID           int
	FirstName    string
	LastName     string
	Birthday     string
	Email        string
	Password     string
	PasswordHash string
}

type Student struct {
	User
	GroupID int
}

type Mentor struct {
	User
	Groups []Group
}

type Login struct {
	Email    string
	Password string
}
