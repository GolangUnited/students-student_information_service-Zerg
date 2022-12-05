package rest

import "github.com/gin-gonic/gin"

// @Summary Create student note
// @Tags student-notes
// @Security ApiToken
// @Description Create new student note in database
// @ID create-student-note
// @Accept  json
// @Produce  json
// @Param input body models.StudentNotes true "student note data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-notes [post]
func (h *Handler) newStudentNote(c *gin.Context) {
}

// @Summary Get student note by id
// @Tags student-notes
// @Security ApiToken
// @Description Get student note from database by ID
// @ID get-student-note-by-id
// @Accept  json
// @Produce  json
// @Param student_note_id path int true "student note ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-notes/{student_note_id} [get]
func (h *Handler) getStudentNoteByID(c *gin.Context) {
}

// @Summary Update student note
// @Tags student-notes
// @Security ApiToken
// @Description Update student note data in database by id
// @ID updatestudent-note-by-id
// @Accept  json
// @Produce  json
// @Param student_note_id path int true "student note ID to update"
// @Param input body models.StudentNotes true "student note  data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-notes/{student_note_id} [put]
func (h *Handler) updateStudentNoteByID(c *gin.Context) {
}

// @Summary Remove student note
// @Tags student-notes
// @Security ApiToken
// @Description Remove student note from database by id
// @ID removestudent-note-by-id
// @Accept  json
// @Produce  json
// @Param student_note_id path int true "student note ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-notes/{student_note_id} [delete]
func (h *Handler) removeStudentNoteByID(c *gin.Context) {
}

// @Summary Get student note  list
// @Tags student-notes
// @Security ApiToken
// @Description Get list of all student note
// @ID student-notes-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.StudentNotes
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-notes [get]
func (h *Handler) StudentNotesList(c *gin.Context) {
}
