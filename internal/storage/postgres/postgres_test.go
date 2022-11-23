package postgres_test

import (
	"testing"
	"zerg-team-student-information-service/internal/storage"
	"zerg-team-student-information-service/internal/storage/postgres"

	"github.com/stretchr/testify/assert"
)

func TestCreateConnectString(t *testing.T) {
	cfg := postgres.PGConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgres",
		Password: "password",
		SSLMode:  "require",
	}
	res := cfg.CreateConnectString()
	exp := "host=localhost port=5432 user=postgres dbname=postgres password=password sslmode=require"

	assert.Equal(t, exp, res, "invalid connection string")
}

func TestPGConfigImplementsDBConfig(t *testing.T) {
	cfg := &postgres.PGConfig{}
	assert.Implements(t, (*storage.DBConfig)(nil), cfg, "PGConfig not implementing DbConfig interface")
}

func TestNewPostgresConnect(t *testing.T) {
	cfg := &postgres.PGConfig{}

	conn, err := postgres.NewPostgresConnect(cfg)
	if err != nil {
		t.Error(err)
	}

	assert.Implements(
		t,
		(*storage.DBConnect)(nil),
		conn,
		"PostgresConnect not implementing DbConnect interface",
	)
}
