package rest

import (
	"net/http"
	"zerg-team-student-information-service/internal/models"

	"github.com/gin-gonic/gin"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User

	if !h.parseBody(c.Request, &user) {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	id, httpCode, err := h.service.CreateUser(user)
	if err != nil {
		c.JSON(httpCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(httpCode, gin.H{
		"message": "User successfully created",
		"user_id": id,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	var login models.User

	if !h.parseBody(c.Request, &login) {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	httpCode, token, err := h.service.GetUser(login)
	if err != nil {
		c.JSON(httpCode, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(httpCode, gin.H{
		"token":   token,
		"message": "OK",
	})
}
