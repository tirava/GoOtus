// Package http implements http server for previewer.
package http

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/preview"
	"gitlab.com/tirava/image-previewer/internal/models"
)

const (
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	maxHeaderBytes = 1 << 20
)

// StartHTTPServer inits routing and starts web listener.
func StartHTTPServer(logger models.Loggerer, conf models.Config, preview preview.Preview, opts entities.ResizeOptions) {
	handlers := newHandlers(logger, conf, preview, opts)
	srv := &http.Server{
		Addr:           conf.ListenHTTP,
		Handler:        handlers.prepareRoutes(),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		defer cancel()
		handlers.logger.Warnf("Signal received: %s", <-shutdown)
		close(handlers.shutdownOthers)

		time.Sleep(time.Millisecond)

		if err := srv.Shutdown(ctx); err != nil {
			handlers.logger.Errorf("Error while shutdown server: %s", err)
		}
	}()

	handlers.logger.Infof("Starting HTTP server at: %s", conf.ListenHTTP)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		handlers.logger.Errorf(err.Error())
		// nolint:gomnd
		os.Exit(1)
	}

	handlers.logger.Infof("Shutdown HTTP server at: %s", conf.ListenHTTP)
}
