package rest

import "github.com/gin-gonic/gin"

// @Summary Add mentor
// @Tags mentors
// @Description Create new group in database
// @ID new-mentor
// @Accept  json
// @Produce  json
// @Param user_id query int true "New mentor ID"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /mentors [post]
func (h *Handler) addMentor(c *gin.Context) {
}

// @Summary Get mentor by id
// @Tags mentors
// @Description Get mentor from database by ID
// @ID get-mentor-by-id
// @Accept  json
// @Produce  json
// @Param mentor_id path int true "Mentor ID to get"
// @Success 200 {object} models.Mentor
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /mentors/{mentor_id} [get]
func (h *Handler) getMentorByID(c *gin.Context) {
}

// @Summary Remove mentor by id
// @Tags mentors
// @Description Remove mentor from database by id
// @ID remove-mentor-by-id
// @Accept  json
// @Produce  json
// @Param mentor_id path int true "Mentor ID to delete"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /mentors/{mentor_id} [delete]
func (h *Handler) removeMentorByID(c *gin.Context) {
}

// @Summary Get mentors list
// @Tags mentors
// @Description Get list of all mentors
// @ID mentors-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Mentor
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /mentors [get]
func (h *Handler) mentorsList(c *gin.Context) {
}
