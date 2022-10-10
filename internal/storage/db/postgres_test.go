package db_test

import (
	"testing"
	"zerg-team-student-information-service/internal/storage/db"

	"github.com/stretchr/testify/assert"
)

func TestCreateConnectString(t *testing.T) {

	cfg := db.PGConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "db",
		Password: "password",
		SSLMode:  true,
	}
	res := cfg.CreateConnectString()
	exp := "host=localhost port=5432 user=postgres dbname=db password=password sslmode=true"

	assert.Equal(t, exp, res, "invalid connection string")

}

func TestPGConfigImplementsDBConfig(t *testing.T) {
	cfg := &db.PGConfig{}
	assert.Implements(t, (*db.DbConfig)(nil), cfg, "PGConfig not implementing DbConfig interface")
}

func TestNewPostgresConnect(t *testing.T) {
	cfg := &db.PGConfig{}

	conn, err := db.NewPostgresConnect(cfg)
	if err != nil {
		t.Error(err)
	}

	assert.Implements(t, (*db.DBConnect)(nil), conn, "PostgresConnect not implementing DbConnect interface")
}
