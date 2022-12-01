package models

type UserData struct {
	UserID    int    `json:"user_id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"bitrhday"`
	Email     string `json:"email"`
}

type User struct {
	UserData
	Password     string `json:"password"`
	passwordHash string
}

func (u *User) GetPasswordHash() string {
	return u.passwordHash
}

func (u *User) SetPasswordHash(passwordHash string) {
	u.passwordHash = passwordHash
}

type Student struct {
	User
	GroupID int `json:"group_id"`
}

type Mentor struct {
	UserData
	Groups []Group
}

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
