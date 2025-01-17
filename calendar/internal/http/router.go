/*
 * HomeWork-10: Calendar extending HTTP methods
 * Created on 14.11.2019 22:18
 * Copyright (c) 2019 - Eugene Klimov
 */

package http

import (
	"context"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/google/uuid"
	"net/http"
	"path"
	"time"
)

type contextKey string

const contextKeyRequestID contextKey = "requestID"

func (h handler) prepareRoutes() http.Handler {
	siteMux := http.NewServeMux()

	h.addPath("GET /hello/*", h.hello)
	h.addPath("GET /get_event", h.getEvent)
	h.addPath("GET /get_user_events", h.getUserEvents)
	h.addPath("POST /create_event", h.createEvent)
	h.addPath("POST /update_event", h.updateEvent)
	h.addPath("POST /delete_event", h.deleteEvent)
	h.addPath("GET /events_for_day", h.eventsForDay)
	h.addPath("GET /events_for_week", h.eventsForWeek)
	h.addPath("GET /events_for_month", h.eventsForMonth)

	siteHandler := h.pathMiddleware(siteMux)
	siteHandler = h.loggerMiddleware(siteHandler)
	siteHandler = h.panicMiddleware(siteHandler)
	return siteHandler
}

func (h handler) addPath(path string, handler http.HandlerFunc) {
	h.handlers[path] = handler
}

func (h handler) pathMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		check := r.Method + " " + r.URL.Path
		for pattern, handlerFunc := range h.handlers {
			if ok, err := path.Match(pattern, check); ok && err == nil {
				handlerFunc(w, r)
				return
			} else if err != nil {
				h.logger.Error("error match router path: %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError),
					http.StatusInternalServerError)
			}
		}
		h.logger.WithFields(loggers.Fields{
			CodeField:  http.StatusNotFound,
			ReqIDField: getRequestID(r.Context()),
		}).Error("RESPONSE")
		http.NotFound(w, r)
	})
}

func (h handler) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.logger.Debug("Middleware 'panic' PASS")
		defer func() {
			if err := recover(); err != nil {
				h.logger.Error("recovered from panic: %s", err)
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
		)).Info("REQUEST START")
		start := time.Now()
		next.ServeHTTP(w, r)
		h.logger.WithFields(loggers.Fields{
			RespTimeField: time.Since(start),
			ReqIDField:    getRequestID(ctx),
		}).Info("REQUEST END")
	})
}

func getRequestID(ctx context.Context) string {
	reqID := ctx.Value(contextKeyRequestID)
	if key, ok := reqID.(string); ok {
		return key
	}
	return ""
}

func requestFields(r *http.Request, args ...string) loggers.Fields {
	fields := make(loggers.Fields)
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
