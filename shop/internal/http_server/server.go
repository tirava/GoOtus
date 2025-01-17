// Package http implements http server.
package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi"
	"github.com/goware/httplog"
	"github.com/jinzhu/gorm"
	"github.com/rs/zerolog"
)

const (
	serviceName = "shop-http"
	logLevel    = "debug"
)

// Server base struct.
type Server struct {
	listen string
	router chi.Router
	logger zerolog.Logger
	db     *gorm.DB
	error  Error
}

// NewServer returns new server instance.
func NewServer(listen string, db *gorm.DB) *Server {
	logger := httplog.NewLogger(serviceName, httplog.Options{
		JSON:     false,
		LogLevel: logLevel,
	})

	return &Server{
		listen: listen,
		logger: logger,
		db:     db,
		error:  newError(&logger),
	}
}

func (s *Server) StartServer() error {
	s.PrepareRouter()

	s.logger.Info().Msgf("Starting HTTP server at: %s", s.listen)

	ctx, cancel := context.WithCancel(context.Background())
	srv := http.Server{Addr: s.listen, Handler: chi.ServerBaseContext(ctx, s.router)}

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	go func() {
		defer cancel()
		s.logger.Warn().Msgf("Signal received: %s", <-shutdown)

		if err := srv.Shutdown(ctx); err != nil {
			s.logger.Error().Msgf("Error while shutdown server: %s", err)
		}
	}()

	err := srv.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		return err
	}

	s.logger.Info().Msgf("Shutdown HTTP server at: %s", s.listen)

	return nil
}
