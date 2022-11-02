package jwt_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"zerg-team-student-information-service/internal/jwt"
	"zerg-team-student-information-service/internal/models"
)

func TestGetUserToken(t *testing.T) {
	user := models.User{
		Email: "test@example.com",
	}

	token, err := jwt.GetUserToken(user)
	assert.NoError(t, err)
	assert.Greater(t, len(token), 10)
}
