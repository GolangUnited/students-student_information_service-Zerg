package rest

import "github.com/gin-gonic/gin"

// @Summary Add mentor
// @Tags mentors
// @Description Create new group in database
// @ID new-mentor
// @Accept  json
// @Produce  json
// @Param input body int true "mentor id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /mentors/new [post]
func (h *Handler) addMentor(c *gin.Context) {
}

// @Summary Get mentor by id
// @Tags mentors
// @Description Get mentor from database by ID
// @ID get-mentor-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "mentor id to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /mentors/:id [get]
func (h *Handler) getMentorByID(c *gin.Context) {
}

// @Summary Remove mentor by id
// @Tags mentors
// @Description Remove mentor from database by id
// @ID remove-mentor-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "mentor id to remove"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /mentors/:id [delete]
func (h *Handler) removeMentorByID(c *gin.Context) {
}
