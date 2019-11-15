/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 22:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"io"
	"net/http"
)

type handler struct {
	handlers map[string]http.HandlerFunc
	calendar calendar.Calendar
	logger   loggers.Logger
}

func newHandlers(calendar calendar.Calendar) *handler {
	return &handler{
		handlers: make(map[string]http.HandlerFunc),
		calendar: calendar,
		logger:   loggers.Logger{}.GetLogger(),
	}
}

func (h handler) hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")
	userID := query.Get("userid")
	eventID := query.Get("eventid")
	if name == "" {
		name = "default name"
	}
	h.logger.WithFields(loggers.Fields{
		CodeField: http.StatusOK,
		IDField:   getRequestID(r.Context()),
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

func (h handler) createEvent(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "create")
}

func (h handler) updateEvent(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "update")
}

func (h handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "delete")
}

func (h handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "events for day")
}

func (h handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "events for week")
}

func (h handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "events for month")
}
