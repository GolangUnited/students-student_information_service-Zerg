package server

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
)

type Server struct {
	httpServer *http.Server
}

func New(handler http.Handler, port string) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + port,
			Handler:        handler,
			MaxHeaderBytes: 1 << 20, // 1048576 byte -> 1024 kbyte -> 1 Mbyte
			ReadTimeout:    10 * time.Second,
			WriteTimeout:   10 * time.Second,
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
				logrus.Fatalf("Error raised while server Listen And Serve:\n%s", err.Error())
			}
		}
	}()
	logrus.Printf("Server started on port %v\n", s.httpServer.Addr)
	<-ctx.Done()

	logrus.Println("Shutting down server gracefully...")

	if err := s.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("error raised while server shutdown:\n%w", err)
	}

	logrus.Println("Shutting down server gracefully... Done")

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
