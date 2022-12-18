package rest

import (
	"database/sql"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"strings"
)

// @Summary Add mentor
// @Tags mentors
// @Security ApiToken
// @Description Create new group in database
// @ID new-mentor
// @Accept  json
// @Produce  json
// @Param user_id query int true "ID of the user who will be granted mentor rights"
// @Success 200
// @Failure 400,404,409 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentors [post]
func (h *Handler) addMentor(c *gin.Context) {
	id, err := strconv.Atoi(c.Request.FormValue("user_id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad mentor id",
		})
		return
	}

	err = h.service.AddMentor(id)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			c.JSON(http.StatusConflict, gin.H{
				"message": "This user is already a mentor",
			})
		} else if strings.Contains(err.Error(), "violates foreign key constraint") {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "User not found",
			})
		} else {
			logrus.Errorf("user adding error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Get mentor by id
// @Tags mentors
// @Security ApiToken
// @Description Get mentor from database by ID
// @ID get-mentor-by-id
// @Accept  json
// @Produce  json
// @Param mentor_id path int true "Mentor ID to get"
// @Success 200 {object} models.Mentor
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentors/{mentor_id} [get]
func (h *Handler) getMentorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad mentor id",
		})
		return
	}

	mentor, err := h.service.GetMentor(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Mentor not found",
			})
		} else {
			logrus.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}
	c.JSON(http.StatusOK, mentor)
}

// @Summary Remove mentor by id
// @Tags mentors
// @Security ApiToken
// @Description Remove mentor from database by id
// @ID remove-mentor-by-id
// @Accept  json
// @Produce  json
// @Param mentor_id path int true "Mentor ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentors/{mentor_id} [delete]
func (h *Handler) removeMentorByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil || id <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Bad mentor id",
		})
		return
	}

	err = h.service.DeleteMentor(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Mentor not found",
			})
		} else {
			logrus.Errorf("mentor deletion error: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}

	c.Status(http.StatusOK)
}

// @Summary Get mentors list
// @Tags mentors
// @Security ApiToken
// @Description Get list of all mentors
// @ID mentors-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Mentor
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /mentors [get]
func (h *Handler) mentorsList(c *gin.Context) {
	mentors, err := h.service.GetMentorsList()
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "Mentors not found",
			})
		} else {
			logrus.Error(err)
			c.JSON(http.StatusInternalServerError, gin.H{
				"message": "Internal Server Error",
			})
		}
		return
	}
	c.JSON(http.StatusOK, mentors)
}
