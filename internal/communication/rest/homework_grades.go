package rest

import "github.com/gin-gonic/gin"

// @Summary Get homework grades by student ID
// @Tags homework grades
// @Security ApiToken
// @Description Get student's homework grades from database by student ID
// @ID get-hw-grades-by-student-id
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID"
// @Success 200 {object} []models.HomeworkGrade
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/hw-grades [get]
func (h *Handler) getHWGradesByStudentID(c *gin.Context) {
}

// @Summary Create new homework grade
// @Tags homework grades
// @Security ApiToken
// @Description Create new homework grade for particular student
// @ID new-hw-grade
// @Accept  json
// @Produce  json
// @Param input body models.HomeworkGrade true "homework mark"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/hw-grades [post]
func (h *Handler) newHWGrade(c *gin.Context) {
}

// @Summary Update homework grade
// @Tags homework grades
// @Security ApiToken
// @Description Update homework grade for particular student
// @ID update-hw-grade
// @Accept  json
// @Produce  json
// @Param input body models.HomeworkGrade true "homework mark"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/hw-grades [patch]
func (h *Handler) updateHWGrade(c *gin.Context) {
}

// @Summary Delete homework grade
// @Tags homework grades
// @Security ApiToken
// @Description Delete homework grade for particular student
// @ID delete-hw-grade
// @Accept  json
// @Produce  json
// @Param homework_grade_id path int true "Homework ID"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/hw-grades [delete]
func (h *Handler) deleteHWGrade(c *gin.Context) {
}
