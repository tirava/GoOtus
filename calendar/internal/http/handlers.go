/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 31.10.2019 22:08
 * Copyright (c) 2019 - Eugene Klimov
 */

package http

import (
	"errors"
	"fmt"
	"github.com/evakom/calendar/internal/domain/calendar"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/json"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/evakom/calendar/internal/urlform"
	"github.com/evakom/calendar/tools"
	"io"
	"net/http"
	"time"
)

// Constants
const (
	ReqIDField    = "request_id"
	HostField     = "host"
	MethodField   = "method"
	URLField      = "url"
	BrowserField  = "browser"
	RemoteField   = "remote"
	QueryField    = "query"
	CodeField     = "response_code"
	RespTimeField = "response_time"
	EventIDField  = "event_id"
	UserIDField   = "user_id"
	DayField      = "day"
	WeekField     = "week"
	MonthField    = "month"
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

	if err := h.calendar.AddEvent(r.Context(), event); err != nil {
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

	eventNew, err := h.calendar.UpdateEventFromEvent(r.Context(), event)
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

	if err := h.calendar.DelEvent(r.Context(), uid); err != nil {
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
	key := urlform.FormWeek
	value := r.URL.Query().Get(key)
	if err := h.getEventsAndSend(key, value, w, r); err != nil {
		h.logger.Debug("[eventsForWeek] error: %s", err)
	}
}

func (h handler) eventsForMonth(w http.ResponseWriter, r *http.Request) {
	key := urlform.FormMonth
	value := r.URL.Query().Get(key)
	if err := h.getEventsAndSend(key, value, w, r); err != nil {
		h.logger.Debug("[eventsForMonth] error: %s", err)
	}
}

func (h handler) getEventsAndSend(key, value string, w http.ResponseWriter, r *http.Request) error {
	var err error
	var events []models.Event
	var fields loggers.Fields

	switch key {
	case urlform.FormEventID:
		events, err = h.calendar.GetAllEventsFilter(r.Context(), models.Event{
			ID: tools.IDString2UUIDorNil(value),
		})
		fields = loggers.Fields{EventIDField: value}
	case urlform.FormUserID:
		events, err = h.calendar.GetAllEventsFilter(r.Context(), models.Event{
			UserID: tools.IDString2UUIDorNil(value),
		})
		fields = loggers.Fields{UserIDField: value}
	case urlform.FormDay:
		events, err = h.calendar.GetAllEventsFilter(r.Context(), models.Event{
			OccursAt: tools.DayString2TimeOrNil(value),
			Duration: 24 * time.Hour,
		})
		fields = loggers.Fields{DayField: value}
	case urlform.FormWeek:
		events, err = h.calendar.GetAllEventsFilter(r.Context(), models.Event{
			OccursAt: tools.DayString2TimeOrNil(value),
			Duration: 24 * time.Hour * 7,
		})
		fields = loggers.Fields{WeekField: value}
	case urlform.FormMonth:
		events, err = h.calendar.GetAllEventsFilter(r.Context(), models.Event{
			OccursAt: tools.DayString2TimeOrNil(value),
			Duration: 24 * time.Hour * 30,
		})
		fields = loggers.Fields{MonthField: value}
	default:
		err = errors.New("invalid key-value in query for getting events")
		fields = loggers.Fields{}
	}

	fields[ReqIDField] = getRequestID(r.Context())

	if err != nil {
		h.logger.WithFields(fields).Error(err.Error())
		h.error.send(w, http.StatusOK, err, fmt.Sprintf("error while get events, %s=%s", key, value))
		return err
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	result, err := json.NewEventResult(events).Encode()
	if err != nil {
		h.logger.WithFields(fields).Error(err.Error())
		h.error.send(w, http.StatusOK, err, fmt.Sprintf("error while encode %s=%s", key, value))
		return err
	}

	if _, err := io.WriteString(w, result); err != nil {
		h.logger.Error("[%s] error write to response writer", key)
		return err
	}

	h.logger.WithFields(loggers.Fields{
		CodeField:  http.StatusOK,
		ReqIDField: getRequestID(r.Context()),
	}).Info("RESPONSE")

	return nil
}

func (h handler) sendResult(events []models.Event, fromHandler string,
	w http.ResponseWriter, r *http.Request) error {

	result, err := json.NewEventResult(events).Encode()
	if err != nil {
		h.logger.WithFields(loggers.Fields{
			CodeField:  http.StatusOK,
			ReqIDField: getRequestID(r.Context()),
		}).Error(err.Error())
		h.error.send(w, http.StatusOK, err, "error while encode events for send result")
		return err
	}

	if _, err := io.WriteString(w, result); err != nil {
		h.logger.Error("[%s] error write to response writer", fromHandler)
		return err
	}

	return nil
}

func (h handler) parseURLFormValues(w http.ResponseWriter, r *http.Request) (models.Event, error) {
	values := make(urlform.Values)
	values[urlform.FormOccursAt] = r.FormValue(urlform.FormOccursAt)
	values[urlform.FormEventID] = r.FormValue(urlform.FormEventID)
	values[urlform.FormSubject] = r.FormValue(urlform.FormSubject)
	values[urlform.FormBody] = r.FormValue(urlform.FormBody)
	values[urlform.FormLocation] = r.FormValue(urlform.FormLocation)
	values[urlform.FormDuration] = r.FormValue(urlform.FormDuration)
	values[urlform.FormUserID] = r.FormValue(urlform.FormUserID)
	values[urlform.FormAlert] = r.FormValue(urlform.FormAlert)

	event, err := values.DecodeEvent()
	if err != nil {
		h.logger.WithFields(loggers.Fields{
			CodeField:  http.StatusBadRequest,
			ReqIDField: getRequestID(r.Context()),
		}).Error(err.Error())
		h.error.send(w, http.StatusBadRequest, err, "error while decode form values")
		return models.Event{}, err
	}

	return event, nil
}
