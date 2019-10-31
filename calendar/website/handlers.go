/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 22:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	"net/http"
	"time"
)

type handler struct {
	logger models.Logger
}

func newHandler() *handler {
	return &handler{
		logger: models.Logger{}.GetLogger(),
	}
}

func (h handler) prepareRoutes() http.Handler {
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/hello", h.helloHandler)
	siteHandler := h.loggerMiddleware(siteMux)
	return siteHandler
}

func (h handler) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.logger.Fields = requestFields(r)
		h.logger.WithFields().Info("REQUEST")
		start := time.Now()
		next.ServeHTTP(w, r)
		h.logger.WithFields().Info("RESPONSE TIME [%s]", time.Since(start))
	})
}

func (h handler) helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	if name == "" {
		name = "default name"
	}
	h.logger.Fields = requestFields(r)
	h.logger.WithFields().Info("RESPONSE CODE [%d]", http.StatusOK)
	fmt.Fprint(w, "Hello, my name is ", name)
}

func requestFields(r *http.Request) models.Fields {
	fields := make(models.Fields)
	fields["host"] = r.Host
	fields["method"] = r.Method
	fields["url"] = r.URL.Path
	//fields["browser"] = r.
	fields["remote"] = r.RemoteAddr
	return fields
}
