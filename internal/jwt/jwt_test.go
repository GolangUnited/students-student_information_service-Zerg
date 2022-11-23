package jwt_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"zerg-team-student-information-service/internal/jwt"
	"zerg-team-student-information-service/internal/models"
)

const testSecret = "test_secret_string"

var testUser = models.User{Email: "test@example.com"}

func TestUserToken(t *testing.T) {
	tokenString, err := jwt.GenerateUserToken(testUser, testSecret)
	assert.NoError(t, err)
	assert.Greater(t, len(tokenString), 5)

	data, err := jwt.GetDataFromToken(tokenString, testSecret)
	assert.NoError(t, err)
	assert.Equal(t, testUser.Email, data.Email)
}
