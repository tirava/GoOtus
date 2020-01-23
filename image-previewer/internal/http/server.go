/*
 * Project: Image Previewer
 * Created on 22.01.2020 21:15
 * Copyright (c) 2020 - Eugene Klimov
 */

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
	fail           = 1
	readTimeout    = 10 * time.Second
	writeTimeout   = 10 * time.Second
	maxHeaderBytes = 1 << 20
)

// StartHTTPServer inits routing and starts web listener.
func StartHTTPServer(logger models.Loggerer, listenHTTP string,
	preview preview.Preview, opts entities.ResizeOptions) {
	handlers := newHandlers(logger, preview, opts)
	srv := &http.Server{
		Addr:           listenHTTP,
		Handler:        handlers.prepareRoutes(),
		ReadTimeout:    readTimeout,
		WriteTimeout:   writeTimeout,
		MaxHeaderBytes: maxHeaderBytes,
	}

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)

	ctx, cancel := context.WithCancel(context.Background())

	defer cancel()

	go func() {
		handlers.logger.Warnf("Signal received: %s", <-shutdown)

		if err := srv.Shutdown(ctx); err != nil {
			handlers.logger.Errorf("Error while shutdown server: %s", err)
		}
	}()

	handlers.logger.Infof("Starting HTTP server at: %s", listenHTTP)

	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		handlers.logger.Errorf(err.Error())
		os.Exit(fail)
	}

	handlers.logger.Infof("Shutdown HTTP server at: %s", listenHTTP)
}
