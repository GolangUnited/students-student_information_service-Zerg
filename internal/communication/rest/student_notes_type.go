package rest

import "github.com/gin-gonic/gin"

// @Summary Create student note type
// @Tags student-note-types
// @Security ApiToken
// @Description Create new student note type in database
// @ID create-student-note-type
// @Accept  json
// @Produce  json
// @Param input body models.StudentNotesType true "student note type data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-note-types [post]
func (h *Handler) newStudentNoteType(c *gin.Context) {
}

// @Summary Get student note type by id
// @Tags student-note-types
// @Security ApiToken
// @Description Get student note type from database by ID
// @ID get-student-note-type-by-id
// @Accept  json
// @Produce  json
// @Param student_note_type_id path int true "student note type ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-note-types/{student_note_type_id} [get]
func (h *Handler) geStudentNoteTypeByID(c *gin.Context) {
}

// @Summary Update student note type
// @Tags student-note-types
// @Security ApiToken
// @Description Update student note type data in database by id
// @ID update-student-note-type-by-id
// @Accept  json
// @Produce  json
// @Param student_note_type_id path int true "student note type ID to update"
// @Param input body models.StudentNotesType true "student note type data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-note-types/{student_note_type_id} [put]
func (h *Handler) updateStudentNoteTypeByID(c *gin.Context) {
}

// @Summary Remove student note type
// @Tags student-note-types
// @Security ApiToken
// @Description Remove student note type from database by id
// @ID remove-student-note-type-by-id
// @Accept  json
// @Produce  json
// @Param student_note_type_id path int true "student note type ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-note-types/{student_note_type_id} [delete]
func (h *Handler) removeStudentNoteTypeByID(c *gin.Context) {
}

// @Summary Get student note type list
// @Tags student-note-types
// @Security ApiToken
// @Description Get list of all student note type
// @ID student-note-types-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.StudentNotesType
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /student-note-types [get]
func (h *Handler) StudentNoteTypesList(c *gin.Context) {
}
