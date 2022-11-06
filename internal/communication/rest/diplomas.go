package rest

import "github.com/gin-gonic/gin"

// @Summary Create diploma
// @Tags diplomas
// @Description Create new diploma in database
// @ID create-diploma
// @Accept  json
// @Produce  json
// @Param input body models.Diploma true "diploma data"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas/new [post]
func (h *Handler) newDiploma(c *gin.Context) {
}

// @Summary Get diploma by id
// @Tags diplomas
// @Description Get diploma from database by ID
// @ID get-diploma-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "diploma id to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas/:id [get]
func (h *Handler) getDiplomaByID(c *gin.Context) {
}

// @Summary Update diploma
// @Tags diplomas
// @Description Update diploma data in database by id
// @ID update-diploma-by-id
// @Accept  json
// @Produce  json
// @Param input body models.Diploma true "diploma data to update"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas/:id [patch]
func (h *Handler) updateDiplomaByID(c *gin.Context) {
}

// @Summary Remove diploma
// @Tags diplomas
// @Description Remove diploma from database by id
// @ID remove-diploma-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "diploma id to remove"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas/:id [delete]
func (h *Handler) removeDiplomaByID(c *gin.Context) {
}
