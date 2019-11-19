/*
 * HomeWork-9: Calendar protobuf preparation
 * Created on 01.11.2019 13:17
 * Copyright (c) 2019 - Eugene Klimov
 */

package website

import (
	"context"
	"errors"
	"fmt"
	"github.com/evakom/calendar/internal/domain/json"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/domain/urlform"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"io"
	"net/http"
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
)

type contextKey string

const contextKeyRequestID contextKey = "requestID"

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

func getRequestID(ctx context.Context) string {
	reqID := ctx.Value(contextKeyRequestID)
	if key, ok := reqID.(string); ok {
		return key
	}
	return ""
}

func (h handler) getEventsAndSend(key, value string, w http.ResponseWriter, r *http.Request) error {
	var err error
	var events []models.Event
	var fields loggers.Fields

	switch key {
	case urlform.FormEventID:
		events, err = h.calendar.GetAllEventsFilter(models.Event{
			ID: tools.IDString2UUIDorNil(value),
		}, 0)
		fields = loggers.Fields{EventIDField: value}
	case urlform.FormUserID:
		events, err = h.calendar.GetAllEventsFilter(models.Event{
			UserID: tools.IDString2UUIDorNil(value),
		}, 0)
		fields = loggers.Fields{UserIDField: value}
	case urlform.FormDay:
		events, err = h.calendar.GetAllEventsFilter(models.Event{
			OccursAt: tools.DayString2TimeOrNil(value),
		}, 1)
		fields = loggers.Fields{DayField: value}
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
