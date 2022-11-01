package repository_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"zerg-team-student-information-service/internal/storage/repository"
)

func TestNew(t *testing.T) {
	repo, err := repository.New("not_valid_dbsm", nil)
	assert.Error(t, err)
	assert.Nil(t, repo)
}
