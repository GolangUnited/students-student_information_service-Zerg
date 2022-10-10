package db

import (
	"database/sql"
	"errors"
)

type DBConnect interface {
	CheckConn() error
	GetConn() *sql.DB
	Close() error
}

type DbConfig interface {
	CreateConnectString() string
}

func NewConnect(dbms string, cfg DbConfig) (DBConnect, error) {
	switch dbms {
	case "postgres":
		db, err := NewPostgresConnect(cfg)
		if err != nil {
			return nil, err
		}
		// if err := db.CheckConn(); err != nil {
		// 	return nil, err
		// }
		return db, nil
	default:
		return nil, errors.New("DBMS not supported")
	}
}
