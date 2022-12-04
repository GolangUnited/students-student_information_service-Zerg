package rest

import "github.com/gin-gonic/gin"

// @Summary Create group
// @Tags groups
// @Security ApiToken
// @Description Create new group in database
// @ID create-group
// @Accept  json
// @Produce  json
// @Param input body models.Group true "group data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /groups [post]
func (h *Handler) newGroup(c *gin.Context) {
}

// @Summary Get group by id
// @Tags groups
// @Security ApiToken
// @Description Get group from database by ID
// @ID get-group-by-id
// @Accept  json
// @Produce  json
// @Param group_id path int true "Group ID to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /groups/{group_id} [get]
func (h *Handler) getGroupByID(c *gin.Context) {
}

// @Summary Update group
// @Tags groups
// @Security ApiToken
// @Description Update group data in database by id
// @ID update-group-by-id
// @Accept  json
// @Produce  json
// @Param group_id path int true "Group ID to update"
// @Param input body models.Group true "group data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /groups/{group_id} [put]
func (h *Handler) updateGroupByID(c *gin.Context) {
}

// @Summary Remove group
// @Tags groups
// @Security ApiToken
// @Description Remove group from database by id
// @ID remove-group-by-id
// @Accept  json
// @Produce  json
// @Param group_id path int true "Group ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /groups/{group_id} [delete]
func (h *Handler) removeGroupByID(c *gin.Context) {
}

// @Summary Get groups list
// @Tags groups
// @Security ApiToken
// @Description Get list of all groups
// @ID groups-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Group
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /groups [get]
func (h *Handler) groupsList(c *gin.Context) {
}
