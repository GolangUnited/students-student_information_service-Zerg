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
		auth.POST("/sign-in", h.signin)
	}

	api := router.Group("/api", h.api)
	{
		users := api.Group("/users")
		{
			users.POST("/", h.newUser)
			users.GET("/:id", h.getUserByID)
			users.PATCH("/:id", h.updateUserByID)
			users.DELETE("/:id", h.removeUserByID)
			users.GET("/", h.usersList)
		}

		students := api.Group("/students")
		{
			students.POST("/", h.addStudent)
			students.GET("/:id", h.getStudentByID)
			students.DELETE("/:id", h.removeStudentByID)
			students.GET("/", h.studentsList)

			items := api.Group(":id")
			{
				items.GET("/certs", h.getCertsByStudentID)
				items.GET("/diploma", h.getDiplomaByStudentID)
				items.GET("/interview", h.getInterviewByStudentID)
				items.GET("/hw-grades", h.getHWGradesByStudentID)
			}
		}

		groups := api.Group("/groups")
		{
			groups.POST("/", h.newGroup)
			groups.GET("/:id", h.getGroupByID)
			groups.PATCH("/:id", h.updateGroupByID)
			groups.DELETE("/:id", h.removeGroupByID)
			groups.GET("/", h.groupsList)
		}

		mentors := api.Group("/mentors")
		{
			mentors.POST("/", h.addMentor)
			mentors.GET("/:id", h.getMentorByID)
			mentors.DELETE("/:id", h.removeMentorByID)
			mentors.GET("/", h.mentorsList)
		}

		admins := api.Group("/admins")
		{
			admins.POST("/", h.addAdmin)
			admins.GET("/:id", h.getAdminByID)
			admins.DELETE("/:id", h.removeAdminByID)
			admins.GET("/", h.adminsList)
		}

		interviews := api.Group("/interviews")
		{
			interviews.POST("/", h.newInterview)
			interviews.GET("/:id", h.getInterviewByID)
			interviews.PATCH("/:id", h.updateInterviewByID)
			interviews.DELETE("/:id", h.removeInterviewByID)
			interviews.GET("/", h.interviewsList)
		}

		diplomas := api.Group("/diplomas")
		{
			diplomas.POST("/", h.newDiploma)
			diplomas.GET("/:id", h.getDiplomaByID)
			diplomas.PATCH("/:id", h.updateDiplomaByID)
			diplomas.DELETE("/:id", h.removeDiplomaByID)
			diplomas.GET("/", h.diplomasList)
		}

		certs := api.Group("/certs")
		{
			certs.POST("/", h.newCert)
			certs.GET("/:id", h.getCertByID)
			certs.PATCH("/:id", h.updateCertByID)
			certs.DELETE("/:id", h.removeCertByID)
			certs.GET("/", h.certsList)
		}

		homework := api.Group("/homework")
		{
			homework.POST("/", h.newHomework)
			homework.GET("/:id", h.getHomeworkByID)
			homework.PATCH("/:id", h.updateHomeworkByID)
			homework.DELETE("/:id", h.removeHomeworkByID)
			homework.GET("/", h.HomeworksList)
		}
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	return router
}
