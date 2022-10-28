package rest

import (
	"database/sql"
	"net/http"
	"zerg-team-student-information-service/internal/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func (h *Handler) CreateUser(c *gin.Context) {
	var user models.User

	if !h.parseBody(c.Request, &user) {
		c.JSON(http.StatusBadRequest, "Неправильный запрос")
		return
	}

	err := user.Validate()
	if err != nil {
		h.logger.Warn(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ex, err := h.service.Repo().User().EmailExist(user.Email)
	if ex {
		c.JSON(http.StatusForbidden, gin.H{
			"message": "Пользователь с таким email-ом существует",
		})
		return
	}
	if err != nil && err != sql.ErrNoRows {
		h.logger.Warn(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), 12)
	if err != nil {
		h.logger.Warn(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}
	user.PasswordHash = string(hash)
	id, err := h.service.Repo().User().Create(user)
	if err != nil {
		h.logger.Warn(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Пользователь был успешно создан",
		"user_id": id,
	})
}

func (h *Handler) GetUser(c *gin.Context) {
	var login models.User

	if !h.parseBody(c.Request, &login) {
		c.JSON(http.StatusBadRequest, "Неправильный запрос")
		return
	}

	err := login.SingUpValidate()
	if err != nil {
		h.logger.Warn(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	user, err := h.service.Repo().User().GetByEmail(login.Email)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Неправильный email или пароль",
		})
		return
	}
	if err != nil {
		h.logger.Warn(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Ошибка сервера",
		})
		return
	}

	hash := []byte(user.PasswordHash)
	err = bcrypt.CompareHashAndPassword(hash, []byte(login.PasswordHash))
	if err != nil {
		h.logger.Warn(err)
		c.JSON(http.StatusUnauthorized, gin.H{
			"message": "Неправильный email или пароль",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
