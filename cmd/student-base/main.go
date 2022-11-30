package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"zerg-team-student-information-service/internal/communication/rest"
	"zerg-team-student-information-service/internal/logger"
	"zerg-team-student-information-service/internal/server"
	"zerg-team-student-information-service/internal/service"
	"zerg-team-student-information-service/internal/storage/repository"
)

// @title           Zerg Team Student Information Service
// @version         0.1
// @description     API for management students database

// @host      localhost:8080

// @securityDefinitions.apikey ApiToken
// @in header
// @name Authorization

func main() {
	appPort := os.Getenv("APP_PORT")
	logrusLogger := logger.NewLogrusLogger()

	if appPort == "" {
		logrusLogger.Warn("APP_PORT environment variable is not set")
		logrusLogger.Warn("Setting port to :8080")
		appPort = "8080"
	}

	repo, err := repository.New("postgres", logrusLogger)
	if err != nil {
		logrusLogger.Error(err)
	}

	svc := service.New(repo, logrusLogger)

	handler := rest.NewHandler(svc, logrusLogger)

	srv := server.New(handler.InitRoutes(), appPort, logrusLogger)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	if err := srv.Run(ctx); err != nil {
		logrusLogger.Error("Server start failed", err)
	}
}
