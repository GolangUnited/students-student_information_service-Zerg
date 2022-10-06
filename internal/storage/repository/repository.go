package repository

import "zerg-team-student-information-service/internal/storage/db"

type Repository struct {
	dbConn db.DBConnect
}

func New(dbConn db.DBConnect) *Repository {
	return &Repository{
		dbConn: dbConn,
	}
}
