/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 18:18
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package website implements http server control.
package website

import (
	"context"
	"github.com/evakom/calendar/internal/configs"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// StartWebsite inits routing and starts web listener.
func StartWebsite(conf configs.Config) {

	handlers := newHandlers()
	srv := &http.Server{
		Addr:    conf.ListenHTTP,
		Handler: handlers.prepareRoutes(),
	}

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		handlers.logger.Info("Signal received: %s", <-shutdown)
		if err := srv.Shutdown(ctx); err != nil {
			handlers.logger.Error("Error while shutdown server:", err)
		}
	}()

	handlers.logger.Info("Starting server at: %s", conf.ListenHTTP)
	handlers.logger.Info("Shutdown server at: %s\n%v", conf.ListenHTTP, srv.ListenAndServe())
}
