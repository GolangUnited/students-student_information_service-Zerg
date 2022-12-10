package models

import (
	"errors"
	"time"
	"zerg-team-student-information-service/internal/service/validator"
)

type UserData struct {
	Login      string `json:"login"`
	FirstName  string `json:"first_name"`
	LastName   string `json:"last_name"`
	Patronymic string `json:"patronymic"`
	Birthday   string `json:"bitrhday"`
	Email      string `json:"email"`
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

func (u *User) Validate() error {
	if len(u.FirstName) == 0 {
		return errors.New("name cannot be empty")
	}

	if len(u.FirstName) > 255 {
		return errors.New("invalid name length")
	}

	if len(u.LastName) == 0 {
		return errors.New("last name cannot be empty")
	}

	if len(u.LastName) > 255 {
		return errors.New("invalid last name length")
	}

	oneYear := 24 * time.Hour * 365

	// TO FIX
	bday, _ := time.Parse(time.RFC822, u.Birthday)

	if time.Since(bday) > 100*oneYear || time.Since(bday) < 9*oneYear {
		return errors.New("invalid date of birth")
	}

	if err := u.LoginAndPasswordValidation(); err != nil {
		return err
	}

	return nil
}

func (u *User) LoginAndPasswordValidation() error {
	if !validator.Matches(u.Email, validator.EmailRX) {
		return errors.New("invalid email")
	}

	if !validator.MinChars(u.Password, 8) {
		return errors.New("password must be at least 8 characters long")
	}

	return nil
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
