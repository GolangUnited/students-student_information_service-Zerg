package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"zerg-team-student-information-service/internal/models"
)

const JwtEnvKey = "JWT_SECRET"

func GetUserToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email": user.Email,
	})

	return token.SignedString([]byte(os.Getenv(JwtEnvKey)))
}

func GetDataFromToken(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv(JwtEnvKey)), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token setup error")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token parsing error")
	}

	return claims, nil
}
