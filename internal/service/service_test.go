package service_test

import (
	"testing"
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/service"
	"zerg-team-student-information-service/internal/storage/repository"

	"github.com/stretchr/testify/assert"
)

func TestNewService(t *testing.T) {
	svc := service.New(&repository.Repository{}, logger.NewLogrusLogger())

	assert.IsType(t, &service.Service{}, svc, "Invalid service type")
}
