package jwt

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"time"
	"zerg-team-student-information-service/internal/models"
)

const JwtEnvKey = "JWT_SECRET"

type TokenData struct {
	Email string
	jwt.RegisteredClaims
}

func GenerateUserToken(user models.User, jwtSecret string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &TokenData{
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	return token.SignedString([]byte(jwtSecret))
}

func GetDataFromToken(tokenString string, jwtSecret string) (*TokenData, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenData{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(jwtSecret), nil
	})
	if err != nil {
		return nil, fmt.Errorf("token parsing error: %w", err)
	}

	data, ok := token.Claims.(*TokenData)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("token parsing error")
	}

	return data, nil
}
