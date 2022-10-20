package models

type User struct {
	Id           int
	FirstName    string
	LastName     string
	Birthday     string
	Email        string
	PasswordHash string
}
