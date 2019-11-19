/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 22:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/domain/urlform"
	"github.com/evakom/calendar/internal/loggers"
	"io"
	"net/http"
)

type handler struct {
	handlers map[string]http.HandlerFunc
	calendar calendar.Calendar
	logger   loggers.Logger
	error    Error
}

func newHandlers(calendar calendar.Calendar) *handler {
	return &handler{
		handlers: make(map[string]http.HandlerFunc),
		calendar: calendar,
		logger:   loggers.GetLogger(),
		error:    newError(),
	}
}

func (h handler) hello(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	name := query.Get("name")

	if name == "" {
		name = "nobody"
	}
	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")

	s := "Hello, my name is " + name + "\n\n"

	if _, err := io.WriteString(w, s); err != nil {
		h.logger.Error("[hello] error write to response writer")
	}
}

func (h handler) getEvent(w http.ResponseWriter, r *http.Request) {
	key := urlform.FormEventID
	value := r.URL.Query().Get(key)
	if err := h.getEventsAndSend(key, value, w, r); err != nil {
		h.logger.Debug("[getEvent] error: %s", err)
	}
}

func (h handler) getUserEvents(w http.ResponseWriter, r *http.Request) {
	key := urlform.FormUserID
	value := r.URL.Query().Get(key)
	if err := h.getEventsAndSend(key, value, w, r); err != nil {
		h.logger.Debug("[getUserEvents] error: %s", err)
	}
}

func (h handler) createEvent(w http.ResponseWriter, r *http.Request) {

	event, err := h.parseURLFormValues(w, r)
	if err != nil {
		h.logger.Debug("[createEvent] error: %s", err)
		return
	}

	if err := h.calendar.AddEvent(event); err != nil {
		h.logger.WithFields(loggers.Fields{
			CodeField:    http.StatusInternalServerError,
			ReqIDField:   getRequestID(r.Context()),
			EventIDField: event.ID.String(),
		}).Error(err.Error())
		h.error.send(w, http.StatusInternalServerError, err,
			"error while adding into calendar, event id="+event.ID.String())
		return
	}

	events := make([]models.Event, 0)
	events = append(events, event)

	if err := h.sendResult(events, "createEvent", w, r); err != nil {
		h.logger.Error("[createEvent] error: %s", err)
	}

	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")
}

func (h handler) updateEvent(w http.ResponseWriter, r *http.Request) {

	event, err := h.parseURLFormValues(w, r)
	if err != nil {
		h.logger.Debug("[updateEvent] error: %s", err)
		return
	}

	eventNew, err := h.calendar.UpdateEventFromEvent(event)
	if err != nil {
		h.logger.WithFields(loggers.Fields{
			CodeField:    http.StatusOK,
			ReqIDField:   getRequestID(r.Context()),
			EventIDField: event.ID.String(),
		}).Error(err.Error())
		h.error.send(w, http.StatusOK, err,
			"error while update calendar, event id="+event.ID.String())
		return
	}

	events := make([]models.Event, 0)
	events = append(events, eventNew)

	if err := h.sendResult(events, "updateEvent", w, r); err != nil {
		h.logger.Error("[updateEvent] error: %s", err)
	}

	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")
}

func (h handler) deleteEvent(w http.ResponseWriter, r *http.Request) {
	value := r.FormValue(urlform.FormEventID)

	uid, err := urlform.DecodeID(value)
	if err != nil {
		h.logger.WithFields(loggers.Fields{
			CodeField:  http.StatusBadRequest,
			ReqIDField: getRequestID(r.Context()),
		}).Error(err.Error())
		h.error.send(w, http.StatusBadRequest, err,
			fmt.Sprintf("error while decode event id=%s", value))
		return
	}

	if err := h.calendar.DelEvent(uid); err != nil {
		h.logger.WithFields(loggers.Fields{
			CodeField:    http.StatusOK,
			ReqIDField:   getRequestID(r.Context()),
			EventIDField: uid.String(),
		}).Error(err.Error())
		h.error.send(w, http.StatusOK, err,
			"error while delete from calendar, event id="+uid.String())
		return
	}

	events := make([]models.Event, 0)
	event := models.Event{}
	event.ID = uid
	events = append(events, event)

	if err := h.sendResult(events, "deleteEvent", w, r); err != nil {
		h.logger.Error("[deleteEvent] error: %s", err)
	}

	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")

}

func (h handler) eventsForDay(w http.ResponseWriter, r *http.Request) {
	key := urlform.FormDay
	value := r.URL.Query().Get(key)
	if err := h.getEventsAndSend(key, value, w, r); err != nil {
		h.logger.Debug("[eventsForDay] error: %s", err)
	}
}

func (h handler) eventsForWeek(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "events for week")
}

func (h handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "events for month")
}
