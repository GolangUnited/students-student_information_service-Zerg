package jwt_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"zerg-team-student-information-service/internal/jwt"
	"zerg-team-student-information-service/internal/models"
)

const testSecret = "test_secret_string"

var testUser = models.User{Email: "test@example.com"}

func TestUserToken(t *testing.T) {
	os.Setenv(jwt.JwtEnvKey, testSecret)
	defer os.Unsetenv(jwt.JwtEnvKey)

	tokenString, err := jwt.GenerateUserToken(testUser)
	assert.NoError(t, err)
	assert.Greater(t, len(tokenString), 5)

	data, err := jwt.GetDataFromToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, testUser.Email, data.Email)
}
