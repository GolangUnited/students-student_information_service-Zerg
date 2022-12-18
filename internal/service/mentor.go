package service

import (
	"zerg-team-student-information-service/internal/models"
)

func (s *Service) GetMentorsList() ([]models.Mentor, error) {
	return s.repo.Mentor.GetAll()
}

func (s *Service) GetMentor(id int) (models.Mentor, error) {
	return s.repo.Mentor.GetByID(id)
}

func (s *Service) DeleteMentor(id int) error {
	return s.repo.Mentor.DeleteByID(id)
}

func (s *Service) AddMentor(id int) error {
	return s.repo.Mentor.Add(id)
}
