package db

import (
	"database/sql"
	"github.com/DATA-DOG/go-sqlmock"
)

type MockConnect struct {
	conn *sql.DB
	Mock sqlmock.Sqlmock
}

func NewMockConnect() (*MockConnect, error) {
	dbConn, mock, err := sqlmock.New()

	return &MockConnect{
		conn: dbConn,
		Mock: mock,
	}, err
}

func (cfg *MockConnect) CreateConnectString() string {
	return ""
}

func (db *MockConnect) CheckConn() error {
	return db.conn.Ping()
}

func (db *MockConnect) GetConn() *sql.DB {
	return db.conn
}

func (db *MockConnect) Close() error {
	return db.conn.Close()
}
