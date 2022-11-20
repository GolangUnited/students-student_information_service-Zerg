package rest

import "github.com/gin-gonic/gin"

// @Summary Create cert
// @Tags certs
// @Description Create new cert in database
// @ID create-cert
// @Accept  json
// @Produce  json
// @Param input body models.Cert true "cert data"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /certs [post]
func (h *Handler) newCert(c *gin.Context) {
}

// @Summary Get cert by id
// @Tags certs
// @Description Get cert from database by ID
// @ID get-cert-by-id
// @Accept  json
// @Produce  json
// @Success 200 {object} models.Cert
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /certs/:id [get]
func (h *Handler) getCertByID(c *gin.Context) {
}

// @Summary Update cert
// @Tags certs
// @Description Update cert data in database by id
// @ID update-cert-by-id
// @Accept  json
// @Produce  json
// @Param input body models.Cert true "cert data to update"
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /certs/:id [patch]
func (h *Handler) updateCertByID(c *gin.Context) {
}

// @Summary Remove cert
// @Tags certs
// @Description Remove cert from database by id
// @ID remove-cert-by-id
// @Accept  json
// @Produce  json
// @Success 200
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /certs/:id [delete]
func (h *Handler) removeCertByID(c *gin.Context) {
}

// @Summary Get certs list
// @Tags certs
// @Description Get list of all certs
// @ID certs-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Cert
// @Failure 400,404 {object} error
// @Failure 500 {object} error
// @Failure default {object} error
// @Router /certs [get]
func (h *Handler) certsList(c *gin.Context) {
}