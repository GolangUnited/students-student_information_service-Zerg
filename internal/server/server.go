package server

import (
	"context"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func New(handler http.Handler, port string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:    ":" + port,
			Handler: handler,
		},
	}
}

func (s *Server) Run(ctx context.Context) error {
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			switch err.Error() {
			case "http: Server closed":
				logrus.Println(err.Error())
			default:
				logrus.Fatalf("Error raised while server Listen And Serve:%s", err.Error())
			}
		}
	}()
	logrus.Printf("Server started on port %v", s.httpServer.Addr)
	<-ctx.Done()

	logrus.Println("Shutting down server gracefully...")

	if err := s.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("error raised while server shutdown:%w", err)
	}

	logrus.Println("Shutting down server gracefully... Done")

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
