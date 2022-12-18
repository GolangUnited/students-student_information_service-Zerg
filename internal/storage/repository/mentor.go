package repository

import "zerg-team-student-information-service/internal/models"

type MentorRepository interface {
	GetAll() ([]models.Mentor, error)
	GetByID(id int) (models.Mentor, error)
	DeleteByID(id int) error
	Add(id int) error
}
