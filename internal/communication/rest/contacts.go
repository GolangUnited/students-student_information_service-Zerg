package rest

import "github.com/gin-gonic/gin"

// @Summary Create contact
// @Tags contacts
// @Security ApiToken
// @Description Create new contact in database
// @ID create-contact
// @Accept  json
// @Produce  json
// @Param input body models.Contact true "contact data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contacts [post]
func (h *Handler) newContact(c *gin.Context) {
}

// @Summary Get contact by id
// @Tags contacts
// @Security ApiToken
// @Description Get contact from database by ID
// @ID get-contact-by-id
// @Accept  json
// @Produce  json
// @Param contact_id path int true "contact ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contacts/{contact_id} [get]
func (h *Handler) getContactByID(c *gin.Context) {
}

// @Summary Update contact
// @Tags contacts
// @Security ApiToken
// @Description Update contact data in database by id
// @ID update-contact-by-id
// @Accept  json
// @Produce  json
// @Param contact_id path int true "contact ID to update"
// @Param input body models.Contact true "contact data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contacts/{contact_id} [put]
func (h *Handler) updateContactByID(c *gin.Context) {
}

// @Summary Remove contact
// @Tags contacts
// @Security ApiToken
// @Description Remove contact from database by id
// @ID remove-contact-by-id
// @Accept  json
// @Produce  json
// @Param contact_id path int true "contact ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contacts/{contact_id} [delete]
func (h *Handler) removeContactByID(c *gin.Context) {
}

// @Summary Get contacts list
// @Tags contacts
// @Security ApiToken
// @Description Get list of all contacts
// @ID contacts-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Contact
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /contacts [get]
func (h *Handler) contactsList(c *gin.Context) {
}
