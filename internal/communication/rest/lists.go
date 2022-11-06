package rest

import "github.com/gin-gonic/gin"

// @Summary Get users list
// @Tags lists
// @Description Get list of all users
// @ID users-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /lists/users [get]
func (h *Handler) usersList(c *gin.Context) {
}

// @Summary Get groups list
// @Tags lists
// @Description Get list of all groups
// @ID groups-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /lists/groups [get]
func (h *Handler) groupsList(c *gin.Context) {
}

// @Summary Get mentors list
// @Tags lists
// @Description Get list of all mentors
// @ID mentors-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /lists/mentors [get]
func (h *Handler) mentorsList(c *gin.Context) {
}

// @Summary Get admins list
// @Tags lists
// @Description Get list of all admins
// @ID admins-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /lists/admins [get]
func (h *Handler) adminsList(c *gin.Context) {
}

// @Summary Get interviews list
// @Tags lists
// @Description Get list of all interviews
// @ID interviews-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /lists/interviews [get]
func (h *Handler) interviewsList(c *gin.Context) {
}

// @Summary Get diplomas list
// @Tags lists
// @Description Get list of all diplomas
// @ID diplomas-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /lists/diplomas [get]
func (h *Handler) diplomasList(c *gin.Context) {
}

// @Summary Get certs list
// @Tags lists
// @Description Get list of all certs
// @ID certs-list
// @Accept  json
// @Produce  json
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /lists/certs [get]
func (h *Handler) certsList(c *gin.Context) {
}
