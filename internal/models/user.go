package models

type UserData struct {
	Login      string    `json:"login"`
	FirstName  string    `json:"first_name"`
	LastName   string    `json:"last_name"`
	Patronymic string    `json:"patronymic"`
	Birthday   string    `json:"bitrhday"`
	Email      string    `json:"email"`
	Contacts   []Contact `json:"contacts"`
}

type User struct {
	UserID int `json:"user_id"`
	UserData
	Password          string `json:"password"`
	PrimaryContacatID int    `json:"primary_contact_id"`
	passwordHash      string
}

func (u *User) GetPasswordHash() string {
	return u.passwordHash
}

func (u *User) SetPasswordHash(passwordHash string) {
	u.passwordHash = passwordHash
}

type StudentData struct {
	UserID  int `json:"user_id"`
	GroupID int `json:"group_id"`
}

type Student struct {
	UserData
	GroupID int `json:"group_id"`
}

type Mentor struct {
	UserData
	Groups []Group `json:"groups"`
}

type Login struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
