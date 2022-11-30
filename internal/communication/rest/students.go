package rest

import "github.com/gin-gonic/gin"

// @Summary Add student
// @Tags students
// @Description Create new group in database
// @ID new-student
// @Accept  json
// @Produce  json
// @Param input body models.Student true "student data"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /students [post]
func (h *Handler) addStudent(c *gin.Context) {
}

// @Summary Get student by id
// @Tags students
// @Description Get student from database by ID
// @ID get-student-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Student
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /students/{student_id} [get]
func (h *Handler) getStudentByID(c *gin.Context) {
}

// @Summary Remove student by id
// @Tags students
// @Description Remove student from database by id
// @ID remove-student-by-id
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /students/{student_id} [delete]
func (h *Handler) removeStudentByID(c *gin.Context) {
}

// @Summary Get students list
// @Tags students
// @Description Get list of all students
// @ID students-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Student
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /students [get]
func (h *Handler) studentsList(c *gin.Context) {
}

// @Summary Get cert by student ID
// @Tags students
// @Description Get student's cert from database by student ID
// @ID get-cert-by-student-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Cert
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /students/{student_id}/cert [get]
func (h *Handler) getCertByStudentID(c *gin.Context) {
}

// @Summary Get diploma by student ID
// @Tags students
// @Description Get student's diploma from database by student ID
// @ID get-diploma-by-student-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Diploma
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /students/{student_id}/diploma [get]
func (h *Handler) getDiplomaByStudentID(c *gin.Context) {
}

// @Summary Get interview by student ID
// @Tags students
// @Description Get student's interview from database by student ID
// @ID get-interview-by-student-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Interview
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /students/{student_id}/interview [get]
func (h *Handler) getInterviewByStudentID(c *gin.Context) {
}

// @Summary Get homework grades by student ID
// @Tags students
// @Description Get student's homework grades from database by student ID
// @ID get-hw-grades-by-student-id
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.HomeworkGrade
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /students/{student_id}/hw-grades [get]
func (h *Handler) getHWGradesByStudentID(c *gin.Context) {
}
