package rest

import "github.com/gin-gonic/gin"

// @Summary Add student
// @Tags students
// @Security ApiToken
// @Description Create new student in database
// @ID new-student
// @Accept  json
// @Produce  json
// @Param input body models.StudentData true "student data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students [post]
func (h *Handler) addStudent(c *gin.Context) {
}

// @Summary Get student by id
// @Tags students
// @Security ApiToken
// @Description Get student from database by ID
// @ID get-student-by-id
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID to get"
// @Success 200 {object} models.Student
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id} [get]
func (h *Handler) getStudentByID(c *gin.Context) {
}

// @Summary Change student's group
// @Tags students
// @Security ApiToken
// @Description Change student's group ID
// @ID change-student-group
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID to patch"
// @Param group_id query int true "Group ID to set"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/group/ [patch]
func (h *Handler) changeStudentGroup(c *gin.Context) {
}

// @Summary Remove student by id
// @Tags students
// @Security ApiToken
// @Description Remove student from database by id
// @ID remove-student-by-id
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id} [delete]
func (h *Handler) removeStudentByID(c *gin.Context) {
}

// @Summary Get students list
// @Tags students
// @Security ApiToken
// @Description Get list of all students
// @ID students-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Student
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students [get]
func (h *Handler) studentsList(c *gin.Context) {
}
