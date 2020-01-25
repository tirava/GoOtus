/*
 * Project: Image Previewer
 * Created on 22.01.2020 21:13
 * Copyright (c) 2020 - Eugene Klimov
 */

package http

import (
	"context"
	"net/http"
	"os"
	"regexp"
	"time"

	"github.com/prometheus/common/log"

	"gitlab.com/tirava/image-previewer/internal/models"

	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/slok/go-http-metrics/middleware"

	"github.com/google/uuid"

	metrics "github.com/slok/go-http-metrics/metrics/prometheus"
)

type contextKey string

const contextKeyRequestID contextKey = "requestID"

func (h handler) prepareRoutes() http.Handler {
	siteMux := http.NewServeMux()

	h.addPath("GET /hello/*", h.helloHandler)
	h.addPath("GET /preview(/.*)", h.previewHandler)

	siteHandler := h.pathMiddleware(siteMux)
	siteHandler = h.loggerMiddleware(siteHandler)
	siteHandler = h.panicMiddleware(siteHandler)

	prometMdlw := middleware.New(middleware.Config{
		Recorder: metrics.NewRecorder(metrics.Config{}),
	})
	hPromet := prometMdlw.Handler("", siteHandler)

	go func() {
		h.logger.Infof("Starting HTTP prometheus exporter at: %s", h.prometPort)
		err := http.ListenAndServe(h.prometPort, promhttp.Handler())

		if err != nil && err != http.ErrServerClosed {
			h.logger.Errorf(err.Error())
			os.Exit(fail)
		}
		// need safely shutdown?
		h.logger.Infof("Shutdown HTTP prometheus exporter at: %s", h.prometPort)
	}()

	return hPromet
}

func (h handler) addPath(regex string, handler http.HandlerFunc) {
	h.handlers[regex] = handler
	cache, err := regexp.Compile(regex)

	if err != nil {
		log.Fatal(err)
	}

	h.cacheHandlers[regex] = cache
}

func (h handler) pathMiddleware(http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		check := r.Method + " " + r.URL.Path
		for pattern, handlerFunc := range h.handlers {
			if h.cacheHandlers[pattern].MatchString(check) {
				handlerFunc(w, r)
				return
			}
		}
		h.logger.WithFields(models.LoggerFields{
			CodeField:  http.StatusNotFound,
			ReqIDField: getRequestID(r.Context()),
		}).Errorf("RESPONSE")
		http.NotFound(w, r)
	})
}

func (h handler) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.logger.Debugf("Middleware 'panic' PASS")
		defer func() {
			if err := recover(); err != nil {
				h.logger.Errorf("recovered from panic: %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}

func (h handler) loggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := assignRequestID(r.Context())
		r = r.WithContext(ctx)
		h.logger.WithFields(requestFields(
			r, ReqIDField, HostField, MethodField, URLField,
			BrowserField, RemoteField, QueryField,
		)).Infof("REQUEST START")
		start := time.Now()
		next.ServeHTTP(w, r)
		h.logger.WithFields(models.LoggerFields{
			RespTimeField: time.Since(start),
			ReqIDField:    getRequestID(ctx),
		}).Infof("REQUEST END")
	})
}

func getRequestID(ctx context.Context) string {
	reqID := ctx.Value(contextKeyRequestID)
	if key, ok := reqID.(string); ok {
		return key
	}

	return ""
}

func requestFields(r *http.Request, args ...string) models.LoggerFields {
	fields := make(models.LoggerFields)

	for _, s := range args {
		switch s {
		case ReqIDField:
			fields[ReqIDField] = getRequestID(r.Context())
		case HostField:
			fields[HostField] = r.Host
		case MethodField:
			fields[MethodField] = r.Method
		case URLField:
			fields[URLField] = r.URL.Path
		case BrowserField:
			fields[BrowserField] = r.Header.Get("User-Agent")
		case RemoteField:
			fields[RemoteField] = r.RemoteAddr
		case QueryField:
			fields[QueryField] = r.URL.RawQuery
		}
	}

	return fields
}

func assignRequestID(ctx context.Context) context.Context {
	reqID := uuid.New()
	return context.WithValue(ctx, contextKeyRequestID, reqID.String())
}
