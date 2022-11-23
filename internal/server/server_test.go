package server_test

import (
	"context"
	"testing"
	"time"
	"zerg-team-student-information-service/internal/communication/rest"
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/server"
	"zerg-team-student-information-service/internal/service"

	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	l := logger.NewLogrusLogger()
	h := rest.NewHandler(&service.Service{}, l, rest.NewMiddleware(""))
	srv := server.New(h.InitRoutes(), "8080", l)

	assert.IsType(t, &server.Server{}, srv, "Invalid server type")
}

func TestNewServer_Run(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	l := logger.NewLogrusLogger()
	h := rest.NewHandler(&service.Service{}, l, rest.NewMiddleware(""))
	srv := server.New(h.InitRoutes(), "8080", l)

	go func() {
		if err := srv.Run(ctx); err != nil {
			t.Error(err)
		}
	}()

	<-ctx.Done()
}

func TestNewServer_Shutdown(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Millisecond)
	defer cancel()

	l := logger.NewLogrusLogger()
	h := rest.NewHandler(&service.Service{}, l, rest.NewMiddleware(""))
	srv := server.New(h.InitRoutes(), "8080", l)

	go func() {
		if err := srv.Run(ctx); err != nil {
			t.Error(err)
		}
	}()

	err := srv.Shutdown(ctx)
	assert.Equal(t, nil, err, "Unexpected error")
}
