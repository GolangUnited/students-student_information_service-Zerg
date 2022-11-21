package rest

import (
	"fmt"
	"time"
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service    *service.Service
	logger     logger.Logger
	middleware *Middleware
}

func NewHandler(service *service.Service, logger logger.Logger, middleware *Middleware) *Handler {
	return &Handler{
		service:    service,
		logger:     logger,
		middleware: middleware,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(gin.LoggerWithFormatter(func(params gin.LogFormatterParams) string {
		return fmt.Sprintf("[%s] - %s %s %s - %d\n",
			params.TimeStamp.Format(time.RFC822),
			params.ClientIP,
			params.Method,
			params.Path,
			params.StatusCode,
		)
	}))

	router.GET("/health", h.healthCheck)
	router.POST("/signup", h.CreateUser)
	router.POST("/signin", h.GetUser)

	//test route for authorization testing
	needAuth := router.Group("/api")
	needAuth.Use(h.middleware.auth)
	needAuth.GET("/health", h.healthCheck)

	return router
}
