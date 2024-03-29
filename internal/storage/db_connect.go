package storage

import (
	"database/sql"
)

type DBConnect interface {
	CheckConn() error
	GetConn() *sql.DB
	Close() error
}

type DBConfig interface {
	CreateConnectString() string
}
