package jwt

import (
	"github.com/golang-jwt/jwt/v4"
	"zerg-team-student-information-service/internal/models"
)

var secret = "secret_string"

func GetUserToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	return token.SignedString([]byte(secret))
}
