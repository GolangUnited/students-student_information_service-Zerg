package repository

import "zerg-team-student-information-service/internal/models"

type UserRepository interface {
	GetById(id int) (models.User, error)
	Create(user models.User) (int64, error)
}
