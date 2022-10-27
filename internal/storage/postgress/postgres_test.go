package postgress_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"zerg-team-student-information-service/internal/storage"
	"zerg-team-student-information-service/internal/storage/postgress"
)

func TestCreateConnectString(t *testing.T) {

	cfg := postgress.PGConfig{
		Host:     "localhost",
		Port:     5432,
		Username: "postgres",
		DBName:   "postgress",
		Password: "password",
		SSLMode:  "require",
	}
	res := cfg.CreateConnectString()
	exp := "host=localhost port=5432 user=postgres dbname=postgress password=password sslmode=require"

	assert.Equal(t, exp, res, "invalid connection string")

}

func TestPGConfigImplementsDBConfig(t *testing.T) {
	cfg := &postgress.PGConfig{}
	assert.Implements(t, (*storage.DbConfig)(nil), cfg, "PGConfig not implementing DbConfig interface")
}

func TestNewPostgresConnect(t *testing.T) {
	cfg := &postgress.PGConfig{}

	conn, err := postgress.NewPostgresConnect(cfg)
	if err != nil {
		t.Error(err)
	}

	assert.Implements(t, (*storage.DBConnect)(nil), conn, "PostgresConnect not implementing DbConnect interface")
}
