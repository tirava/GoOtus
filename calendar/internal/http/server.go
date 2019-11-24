/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 18:18
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package http implements http server control.
package http

import (
	"context"
	"github.com/evakom/calendar/internal/domain/calendar"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

// StartHTTPServer inits routing and starts web listener.
func StartHTTPServer(listenHTTP string, calendar calendar.Calendar) {

	handlers := newHandlers(calendar)
	srv := &http.Server{
		Addr:           listenHTTP,
		Handler:        handlers.prepareRoutes(),
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		handlers.logger.Warn("Signal received: %s", <-shutdown)
		if err := srv.Shutdown(ctx); err != nil {
			handlers.logger.Error("Error while shutdown server: %s", err)
		}
	}()

	handlers.logger.Info("Starting HTTP server at: %s", listenHTTP)
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		handlers.logger.Error(err.Error())
		os.Exit(1)
	}
	handlers.logger.Info("Shutdown HTTP server at: %s", listenHTTP)
}
