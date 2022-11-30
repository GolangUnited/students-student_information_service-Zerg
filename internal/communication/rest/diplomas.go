package rest

import "github.com/gin-gonic/gin"

// @Summary Create diploma
// @Tags diplomas
// @Description Create new diploma in database
// @ID create-diploma
// @Accept  json
// @Produce  json
// @Param input body models.Diploma true "diploma data"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas [post]
func (h *Handler) newDiploma(c *gin.Context) {
}

// @Summary Get diploma by id
// @Tags diplomas
// @Description Get diploma from database by ID
// @ID get-diploma-by-id
// @Accept  json
// @Produce  json
// @Param diploma_id path int true "Diploma ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas/{diploma_id} [get]
func (h *Handler) getDiplomaByID(c *gin.Context) {
}

// @Summary Update diploma
// @Tags diplomas
// @Description Update diploma data in database by id
// @ID update-diploma-by-id
// @Accept  json
// @Produce  json
// @Param diploma_id path int true "Diploma ID to update"
// @Param input body models.Diploma true "diploma data to update"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas/{diploma_id} [patch]
func (h *Handler) updateDiplomaByID(c *gin.Context) {
}

// @Summary Remove diploma
// @Tags diplomas
// @Description Remove diploma from database by id
// @ID remove-diploma-by-id
// @Accept  json
// @Produce  json
// @Param diploma_id path int true "Diploma ID to delete"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas/{diploma_id} [delete]
func (h *Handler) removeDiplomaByID(c *gin.Context) {
}

// @Summary Get diplomas list
// @Tags diplomas
// @Description Get list of all diplomas
// @ID diplomas-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Diploma
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /diplomas [get]
func (h *Handler) diplomasList(c *gin.Context) {
}
