package repository_test

import (
	"testing"
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/storage/db"
	"zerg-team-student-information-service/internal/storage/repository"

	"github.com/stretchr/testify/assert"
)

func TestNewRepository(t *testing.T) {
	r := repository.New(&db.PostgresConnect{}, logger.NewLogrusLogger())

	assert.IsType(t, &repository.Repository{}, r, "unexpected Repository type")
}
