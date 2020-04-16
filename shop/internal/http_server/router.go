package http

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func (s *Server) PrepareRouter() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// for production only instead Logger upper
	// r.Use(httplog.RequestLogger(s.logger))
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))

	r.With(stub).Route("/", func(r chi.Router) {
		r.Get("/", s.root)
		r.Get("/health", s.health)
		r.Get("/version", s.version)
	})

	s.router = r
}
