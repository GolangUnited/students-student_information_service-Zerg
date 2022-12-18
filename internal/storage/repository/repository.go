package repository

import (
	"errors"
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/storage"
	"zerg-team-student-information-service/internal/storage/postgress"
)

type Repository struct {
	dbConn storage.DBConnect
	logger logger.Logger
	user   UserRepository
	Mentor MentorRepository
}

func New(dbms string, logger logger.Logger) (*Repository, error) {
	switch dbms {
	case "postgres":
		cfg := postgress.PGConfig{}
		dbConnect, err := postgress.NewPostgresConnect(&cfg)
		if err != nil {
			return nil, err
		}
		if err = dbConnect.CheckConn(); err != nil {
			return nil, err
		}

		return &Repository{dbConnect, logger,
			postgress.NewUserDB(dbConnect),
			postgress.NewMentorDB(dbConnect)}, nil
	default:
		return nil, errors.New("DBMS not supported")
	}
}

func (r *Repository) User() UserRepository {
	return r.user
}
