package service

import (
	"database/sql"
	"errors"
	"strings"
	"zerg-team-student-information-service/internal/jwt"
	"zerg-team-student-information-service/internal/models"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrUserValidation    = errors.New("user validate error")
	ErrDuplicateKeyRu    = errors.New("pq: повторяющееся значение ключа")
	ErrDuplicateKeyEn    = errors.New("pq: duplicate key value")
	ErrUserAlreadyExists = errors.New("user alredy exists")
	ErrIncorrectPassword = errors.New("incorrect password")
	ErrLoginDoesntExist  = errors.New("login doesn't exist")
	ErrServer            = errors.New("server error")
)

func (s *Service) CreateUser(user models.User) (id int64, err error) {

	err = user.Validate()
	if err != nil {
		s.logger.Warn("create user err ", err.Error())
		err = ErrUserValidation
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		s.logger.Warn("create user err ", err.Error())
		err = ErrServer
		return
	}

	user.PasswordHash = string(hash)
	id, err = s.repo.User().Create(user)
	if err != nil && (strings.Contains(err.Error(), ErrDuplicateKeyEn.Error()) ||
		strings.Contains(err.Error(), ErrDuplicateKeyRu.Error())) {
		s.logger.Warn("create user err ", err.Error())
		err = ErrUserAlreadyExists
		return
	}
	if err != nil {
		s.logger.Warn("create user err ", err.Error())
		err = ErrServer
		return
	}

	return id, nil
}

func (s *Service) SignIn(login models.User) (token string, err error) {

	user, err := s.repo.User().GetByEmail(login.Email)
	if err == sql.ErrNoRows {
		s.logger.Warn("sign in err ", err.Error())
		err = ErrLoginDoesntExist
		return
	}
	if err != nil {
		s.logger.Warn("sign in err ", err.Error())
		err = ErrServer
		return
	}

	hash := []byte(user.PasswordHash)

	err = bcrypt.CompareHashAndPassword(hash, []byte(login.Password))
	if err != nil {
		s.logger.Warn("sign in err ", err.Error())
		err = ErrIncorrectPassword
		return
	}

	token, err = jwt.GenerateUserToken(user, jwt.JwtEnvKey)
	if err != nil {
		s.logger.Warn("sign in err ", err.Error())
		err = ErrServer
		return
	}

	return
}
