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
	"zerg-team-student-information-service/internal/storage/db"
	"zerg-team-student-information-service/internal/storage/repository"
)

func main() {

	port := os.Getenv("APP_PORT")
	logger := logger.NewLogrusLogger()

	if port == "" {
		logger.Warn("APP_PORT environment variable is not set")
		logger.Warn("Setting port to :8080")
		port = "8080"
	}

	cfg := db.PGConfig{}

	dbConn, err := db.NewConnect("postgres", &cfg)
	if err != nil {
		logger.Error("DB connection failed", err)
	}

	repo := repository.New(dbConn, logger)

	svc := service.New(repo, logger)

	handler := rest.NewHandler(svc, logger)

	srv := server.New(handler.InitRoutes(), port, logger)

	ctx, _ := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)

	if err := srv.Run(ctx); err != nil {
		logger.Error("Server start failed", err)
	}
}
