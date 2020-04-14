// Package http implements http server.
package http

import (
	"net/http"

	"github.com/go-chi/chi"
)

// Server base struct.
type Server struct {
	listen string
	router chi.Router
}

// NewServer returns new server instance.
func NewServer(listen string) *Server {
	return &Server{
		listen: listen,
	}
}

func (s *Server) StartServer() error {
	s.PrepareRouter()

	err := http.ListenAndServe(s.listen, s.router)
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}
