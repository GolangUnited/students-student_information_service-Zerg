package rest

import "github.com/gin-gonic/gin"

// @Summary Create interview
// @Tags interviews
// @Description Create new interview in database
// @ID create-interview
// @Accept  json
// @Produce  json
// @Param input body models.Interview true "interview data"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /interviews/new [post]
func (h *Handler) newInterview(c *gin.Context) {
}

// @Summary Get interview by id
// @Tags interviews
// @Description Get interview from database by ID
// @ID get-interview-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "interview id to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /interviews/:id [get]
func (h *Handler) getInterviewByID(c *gin.Context) {
}

// @Summary Update interview
// @Tags interviews
// @Description Update interview data in database by id
// @ID update-interview-by-id
// @Accept  json
// @Produce  json
// @Param input body models.Interview true "interview data to update"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /interviews/:id [patch]
func (h *Handler) updateInterviewByID(c *gin.Context) {
}

// @Summary Remove interview
// @Tags interviews
// @Description Remove interview from database by id
// @ID remove-interview-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "interview id to remove"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /interviews/:id [delete]
func (h *Handler) removeInterviewByID(c *gin.Context) {
}
