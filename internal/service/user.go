package service

import (
	"database/sql"
	"errors"
	"net/http"
	"zerg-team-student-information-service/internal/jwt"
	"zerg-team-student-information-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

func (s *Service) CreateUser(user models.User) (id int64, httpCode int, err error) {
	err = user.Validate()
	if err != nil {
		s.logger.Warn(err)
		httpCode = http.StatusBadRequest
		return
	}

	_, err = s.repo.User().GetByEmail(user.Email)
	if err == nil {
		s.logger.Warn(err)
		err = errors.New("a user with this email already exists")
		httpCode = http.StatusForbidden
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 12)
	if err != nil {
		s.logger.Warn(err)
		err = errors.New("server error")
		httpCode = http.StatusInternalServerError
		return
	}

	user.PasswordHash = string(hash)
	id, err = s.repo.User().Create(user)
	if err != nil {
		s.logger.Warn(err)
		err = errors.New("server error")
		httpCode = http.StatusInternalServerError
		return
	}

	return id, http.StatusCreated, nil
}

func (s *Service) GetUser(login models.User) (httpCode int, token string, err error) {
	err = login.LoginAndPasswordValidation()
	if err != nil {
		s.logger.Warn(err)
		httpCode = http.StatusBadRequest
		return
	}

	user, err := s.repo.User().GetByEmail(login.Email)
	if err == sql.ErrNoRows {
		s.logger.Warn(err)
		err = errors.New("Wrong email or password")
		httpCode = http.StatusUnauthorized
		return
	}
	if err != nil {
		s.logger.Warn(err)
		err = errors.New("Server error")
		httpCode = http.StatusInternalServerError
		return
	}

	hash := []byte(user.PasswordHash)

	err = bcrypt.CompareHashAndPassword(hash, []byte(login.PasswordHash))
	if err != nil {
		s.logger.Warn(err)
		err = errors.New("Wrong email or password")
		httpCode = http.StatusUnauthorized
		return
	}

	token, err = jwt.GenerateUserToken(user)
	if err != nil {
		s.logger.Warn(err)
		err = errors.New("Server error")
		httpCode = http.StatusInternalServerError
		return
	}

	httpCode = http.StatusOK

	return
}
