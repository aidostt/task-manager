package server

import (
	"context"
	"net/http"
	"time"

	"github.com/aidostt/task-manager/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.Config, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:         cfg.HTTP.Host + ":" + cfg.HTTP.Port,
			Handler:      handler,
			ReadTimeout:  10 * time.Second, // добавь
			WriteTimeout: 10 * time.Second, // добавь
			IdleTimeout:  60 * time.Second,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
