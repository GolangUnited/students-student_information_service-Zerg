package rest

import "github.com/gin-gonic/gin"

// @Summary Get certs by student ID
// @Tags student-items
// @Security ApiToken
// @Description Get student's certs from database by student ID
// @ID get-certs-by-student-id
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID"
// @Success 200 {object} []models.Cert
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/certs [get]
func (h *Handler) getCertsByStudentID(c *gin.Context) {
}

// @Summary Get contacts by student ID
// @Tags student-items
// @Security ApiToken
// @Description Get student's contacts from database by student ID
// @ID get-contacts-by-student-id
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID"
// @Success 200 {object} []models.Contact
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/contacts [get]
func (h *Handler) getContactsByStudentID(c *gin.Context) {
}

// @Summary Get student notes by student ID
// @Tags student-items
// @Security ApiToken
// @Description Get student's student notes from database by student ID
// @ID get-student-notes-by-student-id
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID"
// @Success 200 {object} models.StudentNotes
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/student-notes [get]
func (h *Handler) getStudentNotesByStudentID(c *gin.Context) {
}

// @Summary Get mentor notes by mentor ID
// @Tags student-items
// @Security ApiToken
// @Description Get student's mentor notes from database by student ID
// @ID get-mentor-notes-by-student-id
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID"
// @Success 200 {object} models.MentorNotes
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/mentor-notes [get]
func (h *Handler) getMentorNotesByStudentID(c *gin.Context) {
}

// @Summary Get interview by student ID
// @Tags student-items
// @Security ApiToken
// @Description Get student's interview from database by student ID
// @ID get-interview-by-student-id
// @Accept  json
// @Produce  json
// @Param student_id path int true "Student ID"
// @Success 200 {object} models.Interview
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /students/{student_id}/interview [get]
func (h *Handler) getInterviewByStudentID(c *gin.Context) {
}
