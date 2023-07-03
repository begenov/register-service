package server

import (
	"context"
	"net/http"

	"github.com/begenov/register-service/internal/config"
)

type Server struct {
	httpServer *http.Server
}

func NewServer(cfg *config.ServerConfig, handler http.Handler) *Server {
	return &Server{
		httpServer: &http.Server{
			Addr:           ":" + cfg.Port,
			ReadTimeout:    cfg.ReadTimeout,
			WriteTimeout:   cfg.WriteTimeout,
			MaxHeaderBytes: cfg.MaxHeaderMegabytes << 20,
		},
	}
}

func (s *Server) Run() error {
	return s.httpServer.ListenAndServe()
}

func (s *Server) Stop(ctx context.Context) error {
	return s.httpServer.Shutdown(ctx)
}
