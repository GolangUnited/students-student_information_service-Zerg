package rest

import "github.com/gin-gonic/gin"

// @Summary Create cert
// @Tags certs
// @Security ApiToken
// @Description Create new cert in database
// @ID create-cert
// @Accept  json
// @Produce  json
// @Param input body models.Cert true "cert data"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /certs [post]
func (h *Handler) newCert(c *gin.Context) {
}

// @Summary Get cert by id
// @Tags certs
// @Security ApiToken
// @Description Get cert from database by ID
// @ID get-cert-by-id
// @Accept  json
// @Produce  json
// @Param cert_id path int true "Cert ID to get"
// @Success 200 {object} models.Cert
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /certs/{cert_id} [get]
func (h *Handler) getCertByID(c *gin.Context) {
}

// @Summary Update cert
// @Tags certs
// @Security ApiToken
// @Description Update cert data in database by id
// @ID update-cert-by-id
// @Accept  json
// @Produce  json
// @Param cert_id path int true "Cert ID to update"
// @Param input body models.Cert true "cert data to update"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /certs/{cert_id} [put]
func (h *Handler) updateCertByID(c *gin.Context) {
}

// @Summary Remove cert
// @Tags certs
// @Security ApiToken
// @Description Remove cert from database by id
// @ID remove-cert-by-id
// @Accept  json
// @Produce  json
// @Param cert_id path int true "Cert ID to delete"
// @Success 200
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /certs/{cert_id} [delete]
func (h *Handler) removeCertByID(c *gin.Context) {
}

// @Summary Get certs list
// @Tags certs
// @Security ApiToken
// @Description Get list of all certs
// @ID certs-list
// @Accept  json
// @Produce  json
// @Success 200 {object} []models.Cert
// @Failure 400,404 {object} rest.ErrorMessage
// @Failure 500 {object} rest.ErrorMessage
// @Failure default {object} rest.ErrorMessage
// @Router /certs [get]
func (h *Handler) certsList(c *gin.Context) {
}
