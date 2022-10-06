package rest

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func healthCheck(c *gin.Context) {
	logrus.Println("Request received from", c.ClientIP())
	c.JSON(http.StatusOK, gin.H{
		"health": "OK",
	})
}
