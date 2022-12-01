package rest

import (
	"github.com/gin-gonic/gin"
)

// @Summary Create user
// @Tags auth
// @Description Sign up new user
// @ID sign-up
// @Accept  json
// @Produce  json
// @Param input body models.User true "user data"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /auth/sign-up [post]
func (h *Handler) signUp(c *gin.Context) {
	h.newUser(c)
}

// @Summary Get user
// @Tags auth
// @Description Get user from database
// @ID sign-in
// @Accept  json
// @Produce  json
// @Param input body models.Login true "user to get"
// @Success 200 {object} models.JWT
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /auth/sign-in [post]
func (h *Handler) signIn(c *gin.Context) {
}

// @Summary Create user
// @Tags users
// @Security ApiToken
// @Description Create new user in database
// @ID new-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "user data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /users [post]
func (h *Handler) newUser(c *gin.Context) {
}

// @Summary Get user by ID
// @Tags users
// @Security ApiToken
// @Description Get user from database by ID
// @ID get-user-by-id
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID to get"
// @Success 200 {object} models.UserData
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /users/{user_id} [get]
func (h *Handler) getUserByID(c *gin.Context) {
}

// @Summary Update user by id
// @Tags users
// @Security ApiToken
// @Description Update user data in database by id
// @ID update-user-by-id
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID to update"
// @Param input body models.User true "user data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /users/{user_id} [patch]
func (h *Handler) updateUserByID(c *gin.Context) {
}

// @Summary Remove user by id
// @Tags users
// @Security ApiToken
// @Description Remove user from database by id
// @ID remove-user-by-id
// @Accept  json
// @Produce  json
// @Param user_id path int true "User ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /users/{user_id} [delete]
func (h *Handler) removeUserByID(c *gin.Context) {
}

// @Summary Get users list
// @Tags users
// @Security ApiToken
// @Description Get list of all users
// @ID users-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.UserData
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /users [get]
func (h *Handler) usersList(c *gin.Context) {
}
