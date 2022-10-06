package rest_test

import (
	"net/http"
	"net/http/httptest"
	"testing"
	handler "zerg-team-student-information-service/internal/communication/rest"
	"zerg-team-student-information-service/internal/service"

	"github.com/stretchr/testify/require"
)

func TestNewHandler(t *testing.T) {
	h := handler.NewHandler(&service.Service{})

	require.IsType(t, &handler.Handler{}, h)
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

	require.Equal(t, http.StatusOK, res.StatusCode)
}
