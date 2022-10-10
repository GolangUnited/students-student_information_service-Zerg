package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "OK",
	})
}
