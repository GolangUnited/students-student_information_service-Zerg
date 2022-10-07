package repository

import (
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/storage/db"
)

type Repository struct {
	dbConn db.DBConnect
	logger logger.Logger
}

func New(dbConn db.DBConnect, logger logger.Logger) *Repository {
	return &Repository{
		dbConn: dbConn,
		logger: logger,
	}
}
