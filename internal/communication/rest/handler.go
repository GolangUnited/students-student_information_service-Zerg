package rest

import (
	"zerg-team-student-information-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}

}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.GET("/health", healthCheck)

	return router
}
