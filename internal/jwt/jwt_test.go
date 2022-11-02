package jwt_test

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"zerg-team-student-information-service/internal/jwt"
	"zerg-team-student-information-service/internal/models"
)

const testTokenString = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InRlc3RAZXhhbXBsZS5jb20ifQ.1Z5vO0nl_5A7llG_c6wgKeZvmBtGh2LMDmr5CceSdgo"
const testSecret = "test_secret_string"

var testUser = models.User{Email: "test@example.com"}

func TestGetUserToken(t *testing.T) {
	os.Setenv(jwt.JwtEnvKey, testSecret)
	defer os.Unsetenv(jwt.JwtEnvKey)

	want := testTokenString
	got, err := jwt.GetUserToken(testUser)
	assert.NoError(t, err)
	assert.Equal(t, want, got)
}

func TestGetDataFromToken(t *testing.T) {
	os.Setenv(jwt.JwtEnvKey, testSecret)
	defer os.Unsetenv(jwt.JwtEnvKey)

	data, err := jwt.GetDataFromToken(testTokenString)
	assert.NoError(t, err)
	assert.NotNil(t, data)
	assert.Equal(t, testUser.Email, data["email"])
}
