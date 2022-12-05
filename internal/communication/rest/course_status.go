package rest

import "github.com/gin-gonic/gin"

// @Summary Create course status
// @Tags course-statuses
// @Security ApiToken
// @Description Create new course status in database
// @ID create-course-status
// @Accept  json
// @Produce  json
// @Param input body models.CourseStatus true "course status data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /course-statuses [post]
func (h *Handler) newCourseStatus(c *gin.Context) {
}

// @Summary Get course status by id
// @Tags course-statuses
// @Security ApiToken
// @Description Get course status from database by ID
// @ID get-course-status-by-id
// @Accept  json
// @Produce  json
// @Param course_status_id path int true "course status ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /course-statuses/{course_status_id} [get]
func (h *Handler) getCourseStatusByID(c *gin.Context) {
}

// @Summary Update course status
// @Tags course-statuses
// @Security ApiToken
// @Description Update course status data in database by id
// @ID update-course-status-by-id
// @Accept  json
// @Produce  json
// @Param course_status_id path int true "course status ID to update"
// @Param input body models.CourseStatus true "course data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /course-statuses/{course_status_id} [put]
func (h *Handler) updateCourseStatusByID(c *gin.Context) {
}

// @Summary Remove course status
// @Tags course-statuses
// @Security ApiToken
// @Description Remove course status from database by id
// @ID remove-course-status-by-id
// @Accept  json
// @Produce  json
// @Param course_status_id path int true "course status ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /course-statuses/{course_status_id} [delete]
func (h *Handler) removeCourseStatusByID(c *gin.Context) {
}

// @Summary Get course statuses list
// @Tags course-statuses
// @Security ApiToken
// @Description Get list of all course statuses
// @ID course-statuses-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.CourseStatus
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /course-statuses [get]
func (h *Handler) courseStatusesList(c *gin.Context) {
}
