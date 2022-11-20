package rest

import "github.com/gin-gonic/gin"

// @Summary Create homework
// @Tags homeworks
// @Description Create new homework in database
// @ID create-homework
// @Accept  json
// @Produce  json
// @Param input body models.Homework true "homework data"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /homeworks [post]
func (h *Handler) newHomework(c *gin.Context) {
}

// @Summary Get homework by id
// @Tags homeworks
// @Description Get homework from database by ID
// @ID get-homework-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Homework
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /homeworks/:id [get]
func (h *Handler) getHomeworkByID(c *gin.Context) {
}

// @Summary Update homework
// @Tags homeworks
// @Description Update homework data in database by id
// @ID update-homework-by-id
// @Accept  json
// @Produce  json
// @Param input body models.Homework true "homework data to update"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /homeworks/:id [patch]
func (h *Handler) updateHomeworkByID(c *gin.Context) {
}

// @Summary Remove homework
// @Tags homeworks
// @Description Remove homework from database by id
// @ID remove-homework-by-id
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /homeworks/:id [delete]
func (h *Handler) removeHomeworkByID(c *gin.Context) {
}

// @Summary Get homeworks list
// @Tags homeworks
// @Description Get list of all homeworks
// @ID homeworks-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Homework
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /homeworks [get]
func (h *Handler) HomeworksList(c *gin.Context) {
}