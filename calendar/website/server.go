/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 18:18
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package website implements http server control.
package website

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const SERVADDR = ":8080"

// StartWebsite inits routing and starts web listener.
func StartWebsite() {

	handlers := newHandler()
	srv := &http.Server{
		Addr:    SERVADDR,
		Handler: handlers.prepareRoutes(),
	}

	shutdown := make(chan os.Signal)
	signal.Notify(shutdown, os.Interrupt, syscall.SIGTERM)
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go func() {
		log.Println("Signal received:", <-shutdown)
		if err := srv.Shutdown(ctx); err != nil {
			log.Println("Error while shutdown server:", err)
		}
	}()

	fmt.Println("Starting server at:", SERVADDR)
	log.Printf("Shutdown server at: %s\n%v", SERVADDR, srv.ListenAndServe())
}
