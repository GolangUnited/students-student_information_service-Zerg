package rest

import (
	"fmt"
	"time"
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/service"

	_ "zerg-team-student-information-service/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Handler struct {
	service *service.Service
	logger  logger.Logger
}

func NewHandler(service *service.Service, logger logger.Logger) *Handler {
	return &Handler{
		service: service,
		logger:  logger,
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

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.createUser)
		auth.POST("/sign-in", h.getUser)
	}

	api := router.Group("/api", h.api)
	{

		users := api.Group("/users")
		{
			users.POST("/new", h.newUser)
			users.GET("/:id", h.getUserByID)
			users.PATCH("/:id", h.updateUserByID)
			users.DELETE("/:id", h.removeUserByID)
		}

		groups := api.Group("/groups")
		{
			groups.POST("/new", h.newGroup)
			groups.GET("/:id", h.getGroupByID)
			groups.PATCH("/:id", h.updateGroupByID)
			groups.DELETE("/:id", h.removeGroupByID)
		}

		// mentors := api.Group("/mentors")
		// {
		// mentors.POST("/new", h.newGroup)
		// mentors.GET("/:id", h.getGroupByID)
		// mentors.PATCH("/:id", h.updateGroupByID)
		// mentors.DELETE("/:id", h.removeGroupByID)
		// }

		// admins := api.Group("/admins")
		// {
		// admins.POST("/new", h.newGroup)
		// admins.GET("/:id", h.getGroupByID)
		// admins.PATCH("/:id", h.updateGroupByID)
		// admins.DELETE("/:id", h.removeGroupByID)
		// }

		// interviews := api.Group("/mentors")
		// {
		// interviews.POST("/new", h.newGroup)
		// interviews.GET("/:id", h.getGroupByID)
		// interviews.PATCH("/:id", h.updateGroupByID)
		// interviews.DELETE("/:id", h.removeGroupByID)
		// }

		// diplomas := api.Group("/mentors")
		// {
		// diplomas.POST("/new", h.newGroup)
		// diplomas.GET("/:id", h.getGroupByID)
		// diplomas.PATCH("/:id", h.updateGroupByID)
		// diplomas.DELETE("/:id", h.removeGroupByID)
		// }

		// certs := api.Group("/mentors")
		// {
		// certs.POST("/new", h.newGroup)
		// certs.GET("/:id", h.getGroupByID)
		// certs.PATCH("/:id", h.updateGroupByID)
		// certs.DELETE("/:id", h.removeGroupByID)
		// }

		// lists := api.Group("/lists")
		// {
		// lists.GET("/:id", h.getGroupByID)
		// lists.GET("/:id", h.getGroupByID)
		// lists.GET("/:id", h.getGroupByID)
		// lists.GET("/:id", h.getGroupByID)
		// lists.GET("/:id", h.getGroupByID)
		// }
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
