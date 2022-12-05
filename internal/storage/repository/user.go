package repository

import "zerg-team-student-information-service/internal/models"

type UserRepository interface {
	GetByID(id int) (models.User, error)
	Create(user models.User) (int64, error)
	GetByEmail(email string) (models.User, error)
}
