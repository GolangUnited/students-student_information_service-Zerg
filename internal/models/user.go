package models

import (
	"errors"
	"time"
	"zerg-team-student-information-service/internal/service/validator"
)

type User struct {
	ID           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Birthday     time.Time `json:"birthday"`
	Email        string    `json:"email"`
	Password     string    `json:"password"`
	PasswordHash string
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
	if time.Since(u.Birthday) > 100*oneYear || time.Since(u.Birthday) < 9*oneYear {
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
