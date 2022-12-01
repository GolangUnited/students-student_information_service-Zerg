package models

type UserData struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Birthday  string `json:"bitrhday"`
	Email     string `json:"email"`
	Contacts  struct {
		Discord  string `json:"discord"`
		Telegram string `json:"telegram"`
	} `json:"contacts"`
}

type User struct {
	UserID int `json:"user_id"`
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

type StudentData struct {
	UserID  int `json:"user_id"`
	GroupID int `json:"group_id"`
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
