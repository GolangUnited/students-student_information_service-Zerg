package rest

import "github.com/gin-gonic/gin"

// @Summary Create mentor note
// @Tags mentor-notes
// @Security ApiToken
// @Description Create new mentor note in database
// @ID create-mentor-note
// @Accept  json
// @Produce  json
// @Param input body models.MentorNotes true "mentor note data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentor-notes [post]
func (h *Handler) newMentorNote(c *gin.Context) {
}

// @Summary Get mentor note by id
// @Tags mentor-notes
// @Security ApiToken
// @Description Get mentor note from database by ID
// @ID get-mentor-note-by-id
// @Accept  json
// @Produce  json
// @Param mentor_note_id path int true "mentor note ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentor-notes/{mentor_note_id} [get]
func (h *Handler) getMentorNoteByID(c *gin.Context) {
}

// @Summary Update mentor note
// @Tags mentor-notes
// @Security ApiToken
// @Description Update mentor note data in database by id
// @ID updatementor-note-by-id
// @Accept  json
// @Produce  json
// @Param mentor_note_id path int true "mentor note ID to update"
// @Param input body models.MentorNotes true "mentor note  data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentor-notes/{mentor_note_id} [put]
func (h *Handler) updateMentorNoteByID(c *gin.Context) {
}

// @Summary Remove mentor note
// @Tags mentor-notes
// @Security ApiToken
// @Description Remove mentor note from database by id
// @ID removementor-note-by-id
// @Accept  json
// @Produce  json
// @Param mentor_note_id path int true "mentor note ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentor-notes/{mentor_note_id} [delete]
func (h *Handler) removeMentorNoteByID(c *gin.Context) {
}

// @Summary Get mentor note  list
// @Tags mentor-notes
// @Security ApiToken
// @Description Get list of all mentor note
// @ID mentor-notes-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.MentorNotes
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentor-notes [get]
func (h *Handler) mentorNotesList(c *gin.Context) {
}
