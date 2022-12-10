package jwt_test

import (
	"testing"
	"zerg-team-student-information-service/internal/jwt"
	"zerg-team-student-information-service/internal/models"

	"github.com/stretchr/testify/assert"
)

const testSecret = "test_secret_string"

var userData = models.UserData{Email: "test@example.com"}
var testUser = models.User{UserData: userData}

func TestUserToken(t *testing.T) {
	tokenString, err := jwt.GenerateUserToken(testUser, testSecret)
	assert.NoError(t, err)
	assert.Greater(t, len(tokenString), 5)

	data, err := jwt.GetDataFromToken(tokenString, testSecret)
	assert.NoError(t, err)
	assert.Equal(t, testUser.Email, data.Email)
}
