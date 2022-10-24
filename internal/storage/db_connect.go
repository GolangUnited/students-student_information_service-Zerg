package storage

import (
	"database/sql"
)

type DBConnect interface {
	CheckConn() error
	GetConn() *sql.DB
	Close() error
}

type DbConfig interface {
	CreateConnectString() string
}
