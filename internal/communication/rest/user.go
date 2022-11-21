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
	H := make(map[string]any)
	if err != nil {
		H["message"] = err.Error()
	} else {
		H["message"] = "User successfully created"
		H["user_id"] = id
	}

	c.JSON(httpCode, H)
}

func (h *Handler) GetUser(c *gin.Context) {
	var login models.User

	if !h.parseBody(c.Request, &login) {
		c.JSON(http.StatusBadRequest, "Bad Request")
		return
	}

	httpCode, token, err := h.service.GetUser(login)
	H := make(map[string]any)
	if err != nil {
		H["message"] = err.Error()
	} else {
		H["tooken"] = token
		H["message"] = "OK"
	}

	c.JSON(httpCode, H)
}
