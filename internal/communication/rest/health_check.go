package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// healthCheck godoc
// @Tags         health
// @Summary      Application health check
// @Description  Get status of webserver
// @Success      200
// @Router       /health [get]
func (h *Handler) healthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"health": "OK",
	})
}
