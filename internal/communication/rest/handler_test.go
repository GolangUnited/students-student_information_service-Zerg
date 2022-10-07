package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	handler "zerg-team-student-information-service/internal/communication/rest"
	"zerg-team-student-information-service/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	h := handler.NewHandler(&service.Service{})

	assert.IsType(t, &handler.Handler{}, h, "unexpected Handler type")

}

func TestNewHandler_Init(t *testing.T) {
	h := handler.NewHandler(&service.Service{})

	router := h.InitRoutes()

	ts := httptest.NewServer(router)
	defer ts.Close()

	res, err := http.Get(ts.URL + "/health")
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, http.StatusOK, res.StatusCode, "status code is not 200")
}
