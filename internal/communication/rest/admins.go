package rest

import "github.com/gin-gonic/gin"

// @Summary Add admin
// @Tags admins
// @Description Create new group in database
// @ID new-admin
// @Accept  json
// @Produce  json
// @Param input body int true "admin id"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /admins/new [post]
func (h *Handler) addAdmin(c *gin.Context) {
}

// @Summary Get admin by id
// @Tags admins
// @Description Get admin from database by ID
// @ID get-admin-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "admin id to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /admins/:id [get]
func (h *Handler) getAdminByID(c *gin.Context) {
}

// @Summary Remove admin by id
// @Tags admins
// @Description Remove admin from database by id
// @ID remove-admin-by-id
// @Accept  json
// @Produce  json
// @Param input body int true "admin id to remove"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /admins/:id [delete]
func (h *Handler) removeAdminByID(c *gin.Context) {
}
