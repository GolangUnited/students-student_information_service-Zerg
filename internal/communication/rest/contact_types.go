package rest

import "github.com/gin-gonic/gin"

// @Summary Create contact type
// @Tags contact-types
// @Security ApiToken
// @Description Create new contact type in database
// @ID create-contact-type
// @Accept  json
// @Produce  json
// @Param input body models.ContactType true "contact type data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contact-types [post]
func (h *Handler) newContactType(c *gin.Context) {
}

// @Summary Get contact type by id
// @Tags contact-types
// @Security ApiToken
// @Description Get contact type from database by ID
// @ID get-contact-type-by-id
// @Accept  json
// @Produce  json
// @Param contact_type_id path int true "contact type ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contact-types/{contact-type-id} [get]
func (h *Handler) getContactTypeByID(c *gin.Context) {
}

// @Summary Update contact type
// @Tags contact-types
// @Security ApiToken
// @Description Update contact type data in database by id
// @ID update-contact-type-by-id
// @Accept  json
// @Produce  json
// @Param contact_type_id path int true "contact type ID to update"
// @Param input body models.ContactType true "contact type data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contact-types/{contact-type-id} [put]
func (h *Handler) updateContactTypeByID(c *gin.Context) {
}

// @Summary Remove contact type
// @Tags contact-types
// @Security ApiToken
// @Description Remove contact type from database by id
// @ID remove-contact-type-by-id
// @Accept  json
// @Produce  json
// @Param contact_type_id path int true "contact type ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contact-types/{contact-type-id} [delete]
func (h *Handler) removeContactTypeByID(c *gin.Context) {
}

// @Summary Get contact type list
// @Tags contact-types
// @Security ApiToken
// @Description Get list of all contact type
// @ID contact-types-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.ContactType
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contact-types [get]
func (h *Handler) contactTypesList(c *gin.Context) {
}
