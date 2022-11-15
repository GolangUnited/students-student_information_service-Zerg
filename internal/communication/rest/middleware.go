package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zerg-team-student-information-service/internal/jwt"
)

const authorisationHeader = "Authorisation"

type Middleware struct {
	jwtSecret string
}

func NewMiddleware(jwtSecret string) *Middleware {
	return &Middleware{
		jwtSecret: jwtSecret,
	}
}
func (m *Middleware) auth(c *gin.Context) {
	token := c.GetHeader(authorisationHeader)
	if token == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "authorization header is null",
		})
		c.Abort()
		return
	}

	data, err := jwt.GetDataFromToken(token, m.jwtSecret)
	if err != nil || data.Email == "" {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "authorization header is invalid",
		})
		c.Abort()
		return
	}

	c.Next()
}
