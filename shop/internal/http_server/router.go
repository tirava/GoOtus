package http

import (
	chiprometheus "github.com/766b/chi-prometheus"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (s *Server) PrepareRouter() {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	// for production only instead Logger upper
	// r.Use(httplog.RequestLogger(s.logger))
	r.Use(middleware.Recoverer)
	r.Use(render.SetContentType(render.ContentTypeJSON))
	r.Use(chiprometheus.NewMiddleware("shop", 1, 10, 50, 100, 250, 500, 1000, 5000))

	r.Handle("/metrics", promhttp.Handler())

	r.With(stub).Route("/", func(r chi.Router) {
		r.Get("/", s.root)
		r.Get("/health", s.health)
		r.Get("/version", s.version)
		r.Get("/users", s.getUsers)
		r.Route("/user", func(r chi.Router) {
			r.Post("/", s.newUser)
			r.Get("/{userID}", s.getUser)
			r.Put("/{userID}", s.updateUser)
			r.Delete("/{userID}", s.deleteUser)
		})
	})

	s.router = r
}
