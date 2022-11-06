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

		mentors := api.Group("/mentors")
		{
			mentors.POST("/new", h.addMentor)
			mentors.GET("/:id", h.getMentorByID)
			mentors.DELETE("/:id", h.removeMentorByID)
		}

		admins := api.Group("/admins")
		{
			admins.POST("/new", h.addAdmin)
			admins.GET("/:id", h.getAdminByID)
			admins.DELETE("/:id", h.removeAdminByID)
		}

		interviews := api.Group("/interviews")
		{
			interviews.POST("/new", h.newInterview)
			interviews.GET("/:id", h.getInterviewByID)
			interviews.PATCH("/:id", h.updateInterviewByID)
			interviews.DELETE("/:id", h.removeInterviewByID)
		}

		diplomas := api.Group("/diplomas")
		{
			diplomas.POST("/new", h.newDiploma)
			diplomas.GET("/:id", h.getDiplomaByID)
			diplomas.PATCH("/:id", h.updateDiplomaByID)
			diplomas.DELETE("/:id", h.removeDiplomaByID)
		}

		certs := api.Group("/certs")
		{
			certs.POST("/new", h.newCert)
			certs.GET("/:id", h.getCertByID)
			certs.PATCH("/:id", h.updateCertByID)
			certs.DELETE("/:id", h.removeCertByID)
		}

		lists := api.Group("/lists")
		{
			lists.GET("/users", h.usersList)
			lists.GET("/groups", h.groupsList)
			lists.GET("/mentors", h.mentorsList)
			lists.GET("/admins", h.adminsList)
			lists.GET("/interviews", h.interviewsList)
			lists.GET("/diplomas", h.diplomasList)
			lists.GET("/certs", h.certsList)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
