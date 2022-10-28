package models

import (
	"errors"
	"time"
	"zerg-team-student-information-service/internal/validator"
)

type User struct {
	Id           int       `json:"id"`
	FirstName    string    `json:"first_name"`
	LastName     string    `json:"last_name"`
	Birthday     time.Time `json:"birthday"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"pass_hash"`
}

func (u *User) Validate() error {
	if u.Id < 0 || u.Id > 2147483647 {
		return errors.New("Неверный id")
	}

	if len(u.FirstName) == 0 {
		return errors.New("Имя не может быть пустым")
	}

	if len(u.FirstName) > 255 {
		return errors.New("Недопустимая длина имени")
	}

	if len(u.LastName) == 0 {
		return errors.New("Фамилия не может быть пустым")
	}

	if len(u.LastName) > 255 {
		return errors.New("Недопустимая длина фамилии")
	}

	oneYear := 24 * time.Hour * 365
	if time.Since(u.Birthday) > 100*oneYear || time.Since(u.Birthday) < 9*oneYear {
		return errors.New("Неверная дата рождения")
	}

	if !validator.Matches(u.Email, validator.EmailRX) {
		return errors.New("Неправильный email")
	}

	if !validator.MinChars(u.PasswordHash, 8) {
		return errors.New("Пароль состоит минимум из 8 символов")
	}

	return nil
}

func (u *User) SingUpValidate() error {
	if !validator.Matches(u.Email, validator.EmailRX) {
		return errors.New("Неправильный email")
	}

	if !validator.MinChars(u.PasswordHash, 8) {
		return errors.New("Пароль состоит минимум из 8 символов")
	}

	return nil
}
