package rest

import "github.com/gin-gonic/gin"

// @Summary Create interview
// @Tags interviews
// @Security ApiToken
// @Description Create new interview in database
// @ID create-interview
// @Accept  json
// @Produce  json
// @Param input body models.Interview true "interview data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /interviews [post]
func (h *Handler) newInterview(c *gin.Context) {
}

// @Summary Get interview by id
// @Tags interviews
// @Security ApiToken
// @Description Get interview from database by ID
// @ID get-interview-by-id
// @Accept  json
// @Produce  json
// @Param interview_id path int true "Interview ID to get"
// @Success 200 {object} models.Interview
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /interviews/{interview_id} [get]
func (h *Handler) getInterviewByID(c *gin.Context) {
}

// @Summary Update interview
// @Tags interviews
// @Security ApiToken
// @Description Update interview data in database by id
// @ID update-interview-by-id
// @Accept  json
// @Produce  json
// @Param interview_id path int true "Interview ID to update"
// @Param input body models.Interview true "interview data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /interviews/{interview_id} [patch]
func (h *Handler) updateInterviewByID(c *gin.Context) {
}

// @Summary Remove interview
// @Tags interviews
// @Security ApiToken
// @Description Remove interview from database by id
// @ID remove-interview-by-id
// @Accept  json
// @Produce  json
// @Param interview_id path int true "Interview ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /interviews/{interview_id} [delete]
func (h *Handler) removeInterviewByID(c *gin.Context) {
}

// @Summary Get interviews list
// @Tags interviews
// @Security ApiToken
// @Description Get list of all interviews
// @ID interviews-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Interview
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /interviews [get]
func (h *Handler) interviewsList(c *gin.Context) {
}
