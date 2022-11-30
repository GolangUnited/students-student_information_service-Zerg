package rest

import (
	"github.com/gin-gonic/gin"
)

// @Summary Get user
// @Tags auth
// @Description Get user from database
// @ID get-user
// @Accept  json
// @Produce  json
// @Param input body models.Login true "user to get"
// @Success 200 {object} models.JWT
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /auth/signin [post]
func (h *Handler) signin(c *gin.Context) {
}

// @Summary Create user
// @Tags users
// @Description Create new user in database
// @ID new-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "user data"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /users [post]
func (h *Handler) newUser(c *gin.Context) {
}

// @Summary Get user by ID
// @Tags users
// @Description Get user from database by ID
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.UserData
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /users/{user_id} [get]
func (h *Handler) getUserByID(c *gin.Context) {
}

// @Summary Update user by id
// @Tags users
// @Description Update user data in database by id
// @ID update-user-by-id
// @Accept  json
// @Produce  json
// @Param input body models.User true "user data to update"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /users/{user_id} [patch]
func (h *Handler) updateUserByID(c *gin.Context) {
}

// @Summary Remove user by id
// @Tags users
// @Description Remove user from database by id
// @ID remove-user-by-id
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /users/{user_id} [delete]
func (h *Handler) removeUserByID(c *gin.Context) {
}

// @Summary Get users list
// @Tags users
// @Description Get list of all users
// @ID users-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.UserData
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /users [get]
func (h *Handler) usersList(c *gin.Context) {
}
