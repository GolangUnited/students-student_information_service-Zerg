package rest

import "github.com/gin-gonic/gin"

// @Summary Add admin
// @Tags admins
// @Description Create new group in database
// @ID new-admin
// @Accept  json
// @Produce  json
// @Param id header int true "New admin ID"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /admins [post]
func (h *Handler) addAdmin(c *gin.Context) {
}

// @Summary Get admin by id
// @Tags admins
// @Description Get admin from database by ID
// @ID get-admin-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserData
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /admins/{admin_id} [get]
func (h *Handler) getAdminByID(c *gin.Context) {
}

// @Summary Remove admin by id
// @Tags admins
// @Description Remove admin from database by id
// @ID remove-admin-by-id
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /admins/{admin_id} [delete]
func (h *Handler) removeAdminByID(c *gin.Context) {
}

// @Summary Get admins list
// @Tags admins
// @Description Get list of all admins
// @ID admins-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.UserData
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /admins [get]
func (h *Handler) adminsList(c *gin.Context) {
}
