package rest

import "github.com/gin-gonic/gin"

// @Summary Create group contact
// @Tags group-contacts
// @Security ApiToken
// @Description Create new group contact in database
// @ID create-group-contact
// @Accept  json
// @Produce  json
// @Param input body models.GroupContact true "group contact data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /group-contacts [post]
func (h *Handler) newGroupContact(c *gin.Context) {
}

// @Summary Get group contact by id
// @Tags group-contacts
// @Security ApiToken
// @Description Get group contact from database by ID
// @ID get-group-contact-by-id
// @Accept  json
// @Produce  json
// @Param group_contact_id path int true "group contact ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /group-contacts/{group_contact_id} [get]
func (h *Handler) getGroupContactByID(c *gin.Context) {
}

// @Summary Update group contact
// @Tags group-contacts
// @Security ApiToken
// @Description Update group contact data in database by id
// @ID update-group-contact-by-id
// @Accept  json
// @Produce  json
// @Param group_contact_id path int true "contact ID to update"
// @Param input body models.GroupContact true "group contact data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /group-contacts/{group_contact_id} [put]
func (h *Handler) updateGroupContactByID(c *gin.Context) {
}

// @Summary Remove group contact
// @Tags group-contacts
// @Security ApiToken
// @Description Remove group contact from database by id
// @ID remove-group-contact-by-id
// @Accept  json
// @Produce  json
// @Param group_contact_id path int true "group contact ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /group-contacts/{group_contact_id} [delete]
func (h *Handler) removeGroupContactByID(c *gin.Context) {
}

// @Summary Get group contacts list
// @Tags group-contacts
// @Security ApiToken
// @Description Get list of all group contacts
// @ID group-contacts-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.GroupContact
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /group-contacts [get]
func (h *Handler) groupContactsList(c *gin.Context) {
}
