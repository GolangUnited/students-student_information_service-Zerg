package service

import (
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/storage/repository"
)

type Service struct {
	repo   *repository.Repository
	logger logger.Logger
}

func New(repo *repository.Repository, logger logger.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: logger,
	}
}

func (s *Service) Repo() *repository.Repository {
	return s.repo
}
