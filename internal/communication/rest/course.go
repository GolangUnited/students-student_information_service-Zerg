package rest

import "github.com/gin-gonic/gin"

// @Summary Create course
// @Tags courses
// @Security ApiToken
// @Description Create new course in database
// @ID create-course
// @Accept  json
// @Produce  json
// @Param input body models.Course true "course data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /courses [post]
func (h *Handler) newCourse(c *gin.Context) {
}

// @Summary Get course by id
// @Tags courses
// @Security ApiToken
// @Description Get course from database by ID
// @ID get-course-by-id
// @Accept  json
// @Produce  json
// @Param course_id path int true "course ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /courses/{course_id} [get]
func (h *Handler) getCourseByID(c *gin.Context) {
}

// @Summary Update course
// @Tags courses
// @Security ApiToken
// @Description Update course data in database by id
// @ID update-course-by-id
// @Accept  json
// @Produce  json
// @Param course_id path int true "course ID to update"
// @Param input body models.Course true "course data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /courses/{course_id} [put]
func (h *Handler) updateCourseByID(c *gin.Context) {
}

// @Summary Remove course
// @Tags courses
// @Security ApiToken
// @Description Remove course from database by id
// @ID remove-course-by-id
// @Accept  json
// @Produce  json
// @Param course_id path int true "course ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /courses/{course_id} [delete]
func (h *Handler) removeCourseByID(c *gin.Context) {
}

// @Summary Get courses list
// @Tags courses
// @Security ApiToken
// @Description Get list of all courses
// @ID courses-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Course
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /courses [get]
func (h *Handler) coursesList(c *gin.Context) {
}

// @Summary Get course interview
// @Tags courses
// @Security ApiToken
// @Description Get course interview from database by course ID
// @ID get-course-interview-by-id
// @Accept  json
// @Produce  json
// @Param course_id path int true "course ID to get interview"
// @Success 200 {object} models.CourseInterview
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /courses/{course_id}/interview [get]
func (h *Handler) getCourseInterviewByID(c *gin.Context) {
}
