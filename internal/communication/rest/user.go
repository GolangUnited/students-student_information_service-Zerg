package rest

import (
	"net/http"
	"zerg-team-student-information-service/internal/models"
	"zerg-team-student-information-service/internal/service"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User

	if !h.parseBody(c.Request, &user) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	id, err := h.service.CreateUser(user)
	switch err {
	case service.ErrUserValidation:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "validation error",
		})
	case service.ErrUserAlreadyExists:
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "User with this email already exists",
		})
	case service.ErrServer:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	default:
		c.JSON(http.StatusCreated, gin.H{
			"message": "User successfully created",
			"user_id": id,
		})
	}
}

func (h *Handler) SignIn(c *gin.Context) {
	var user models.User

	if !h.parseBody(c.Request, &user) {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad Request",
		})
		return
	}

	token, err := h.service.SignIn(user)
	switch err {
	case service.ErrLoginDoesntExist:
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "User with this email doesn't exist",
		})
	case service.ErrIncorrectPassword:
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Incorrect password",
		})
	case service.ErrServer:
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Internal server error",
		})
	default:
		c.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}
}
