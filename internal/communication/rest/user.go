package rest

import (
	"github.com/gin-gonic/gin"
)

// @Summary Create user
// @Tags auth
// @Description Create new user in database
// @ID create-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "user data"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /auth/signup [post]
func (h *Handler) CreateUser(c *gin.Context) {
}

// @Summary Get user
// @Tags auth
// @Description Get user from database
// @ID get-user
// @Accept  json
// @Produce  json
// @Param input body models.User true "user to get"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /auth/signin [post]
func (h *Handler) GetUser(c *gin.Context) {
}
