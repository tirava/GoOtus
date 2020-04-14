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
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(middleware.Heartbeat("/ping"))

	r.With(stub).Route("/", func(r chi.Router) {
		r.Get("/", s.root)
		r.Get("/health", s.health)
	})

	s.router = r
}
