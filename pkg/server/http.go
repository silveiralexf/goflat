package server

import (
	"context"
	"net/http"
	"time"

	"github.com/silveiralexf/goflat/pkg/ui/pages"
)

type Server struct {
	server http.Server
}

func (s *Server) Start() error {
	return s.server.ListenAndServe()
}

func (s *Server) Shutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func New(address string, timeout time.Duration) *Server {
	return &Server{
		server: http.Server{
			Addr:    address,
			Handler: pages.GetHandlerSet(timeout),
		},
	}
}
