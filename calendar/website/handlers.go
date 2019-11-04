/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 22:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"io"
	"net/http"
	"time"
)

type handler struct {
	calendar calendar.Calendar
	logger   models.Logger
}

func newHandlers(calendar calendar.Calendar) *handler {
	return &handler{
		calendar: calendar,
		logger:   models.Logger{}.GetLogger(),
	}
}

func (h handler) prepareRoutes() http.Handler {
	siteMux := http.NewServeMux()
	siteMux.HandleFunc("/hello/", h.helloHandler)
	siteMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		h.logger.WithFields(models.Fields{
			"code": http.StatusNotFound,
			ID:     getRequestID(r.Context()),
		}).Error("RESPONSE")
		http.NotFound(w, r)
	})
	siteHandler := h.loggerMiddleware(siteMux)
	siteHandler = h.panicMiddleware(siteHandler)
	return siteHandler
}

func (h handler) panicMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		h.logger.Debug("Middleware 'panic' PASS")
		defer func() {
			if err := recover(); err != nil {
				h.logger.Error("recovered from panic: %s", err)
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
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
			r, ID, HOSTFIELD, METHODFIELD, URLFIELD,
			BROWSERFIELD, REMOTEFIELD, QUERYFIELD,
		)).Info("REQUEST START")
		start := time.Now()
		next.ServeHTTP(w, r)
		h.logger.WithFields(models.Fields{
			"response_time": time.Since(start),
			ID:              getRequestID(ctx),
		}).Info("REQUEST END")
	})
}

func (h handler) helloHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	userID := query.Get("userid")
	eventID := query.Get("eventid")
	if name == "" {
		name = "default name"
	}
	h.logger.WithFields(models.Fields{
		"code": http.StatusOK,
		ID:     getRequestID(r.Context()),
	}).Info("RESPONSE")

	event := models.NewEvent()
	event.Location = "qqqqqqqqqqqqqqqqqqqqqq"
	event.UserID = uuid.New()
	_ = h.calendar.AddEvent(event)

	s := "Hello, my name is " + name + "\n\n"

	events, err := h.calendar.GetAllEventsFilter(models.Event{
		ID:     tools.IDString2UUIDorNil(eventID),
		UserID: tools.IDString2UUIDorNil(userID),
	})

	if err != nil {
		s += err.Error()
	}

	for _, e := range events {
		s += e.StringEr() + "\n"
	}

	if _, err := io.WriteString(w, s); err != nil {
		h.logger.Error("Error write to response writer!")
	}
}
