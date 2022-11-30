package jwt_test

import (
	"os"
	"testing"
	"zerg-team-student-information-service/internal/jwt"
	"zerg-team-student-information-service/internal/models"

	"github.com/stretchr/testify/assert"
)

const testSecret = "test_secret_string"

var userData = models.UserData{Email: "test@example.com"}
var testUser = models.User{UserData: userData}

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
