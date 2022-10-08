package server

import (
	"context"
	"fmt"
	"net/http"
	"zerg-team-student-information-service/internal/logger"
)

type Server struct {
	httpServer *http.Server
	logger     logger.Logger
}

func New(handler http.Handler, port string, logger logger.Logger) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
		logger: logger,
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			switch err.Error() {
			case "http: Server closed":
				s.logger.Info(err.Error())
			default:
				s.logger.Errorf("Error raised while server Listen And Serve:%s", err.Error())
			}
		}
	}()
	s.logger.Infof("Server started on port %v", s.httpServer.Addr)
	<-ctx.Done()

	s.logger.Info("Shutting down server gracefully...")

	if err := s.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("error raised while server shutdown:%w", err)
	}

	s.logger.Info("Shutting down server gracefully... Done")

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
