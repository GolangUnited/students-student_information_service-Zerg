package rest

import "github.com/gin-gonic/gin"

// @Summary Create group
// @Tags groups
// @Description Create new group in database
// @ID create-group
// @Accept  json
// @Produce  json
// @Param input body models.Group true "group data"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /groups/new [post]
func (h *Handler) newGroup(c *gin.Context) {
	h.createUser(c)
}

// @Summary Get group by it
// @Tags groups
// @Description Get group from database by ID
// @ID get-group-by-id
// @Accept  json
// @Produce  json
// @Param input body models.Group true "group id to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /groups/:id [get]
func (h *Handler) getGroupByID(c *gin.Context) {
	h.createUser(c)
}

// @Summary Update group
// @Tags groups
// @Description Update group data in database by id
// @ID update-group-by-id
// @Accept  json
// @Produce  json
// @Param input body models.Group true "group data to update"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /groups/:id [patch]
func (h *Handler) updateGroupByID(c *gin.Context) {
	h.createUser(c)
}

// @Summary Delete user
// @Tags groups
// @Description Remove group from database by id
// @ID delete-group-by-id
// @Accept  json
// @Produce  json
// @Param input body models.Group true "group id to delete"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /groups/:id [delete]
func (h *Handler) removeGroupByID(c *gin.Context) {
	h.createUser(c)
}
