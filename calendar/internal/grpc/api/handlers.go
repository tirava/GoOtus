/*
 * HomeWork-12: gRPC server
 * Created on 24.11.2019 14:40
 * Copyright (c) 2019 - Eugene Klimov
 */

package api

import (
	"context"
	"github.com/evakom/calendar/internal/domain/errors"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/golang/protobuf/ptypes"
	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Constants
const (
	CodeField    = "response_code"
	EventIDField = "event_id"
)

// CreateEvent creates event.
func (cs *CalendarServer) CreateEvent(ctx context.Context, req *EventRequest) (*EventResponse, error) {

	cs.logger.Info("REQUEST [CreateEvent]")

	event := models.NewEvent()

	protoEvent := &Event{
		Id:       event.ID.String(),
		OccursAt: req.GetOccursAt(),
		Subject:  req.GetSubject(),
		Body:     req.GetBody(),
		Duration: req.GetDuration(),
		Location: req.GetLocation(),
		UserID:   req.GetUserID(),
	}

	occursAt, err := ptypes.Timestamp(protoEvent.OccursAt)
	if err != nil {
		cs.logger.WithFields(loggers.Fields{
			CodeField: codes.InvalidArgument,
		}).Error("RESPONSE [CreateEvent]: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	duration, err := ptypes.Duration(protoEvent.Duration)
	if err != nil {
		cs.logger.WithFields(loggers.Fields{
			CodeField: codes.InvalidArgument,
		}).Error("RESPONSE [CreateEvent]: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	uid, err := uuid.Parse(protoEvent.UserID)
	if err != nil {
		cs.logger.WithFields(loggers.Fields{
			CodeField: codes.InvalidArgument,
		}).Error("RESPONSE [CreateEvent]: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	event.OccursAt = occursAt
	event.Subject = protoEvent.Subject
	event.Body = protoEvent.Body
	event.Duration = duration
	event.Location = protoEvent.Location
	event.UserID = uid

	if err := cs.calendar.AddEvent(event); err != nil {
		cs.logger.WithFields(loggers.Fields{
			CodeField: codes.Internal,
		}).Error("RESPONSE [CreateEvent]: %s", err)
		if bizErr, ok := err.(errors.EventError); ok {
			resp := &EventResponse{
				Result: &EventResponse_Error{
					Error: bizErr.Error(),
				},
			}
			return resp, nil
		}
		return nil, status.Error(codes.Internal, err.Error())
	}

	cs.logger.WithFields(loggers.Fields{
		CodeField:    codes.OK,
		EventIDField: protoEvent.Id,
	}).Info("RESPONSE [CreateEvent]")

	resp := &EventResponse{
		Result: &EventResponse_Event{
			Event: protoEvent,
		},
	}
	cs.logger.Debug("[CreateEvent] Response body: %+v", resp)

	return resp, nil
}

// GetEvent got one event by id.
func (cs *CalendarServer) GetEvent(context.Context, *ID) (*EventResponse, error) {
	panic("GetEvent implement me")
}

// GetUserEvents returns all events for given user.
func (cs *CalendarServer) GetUserEvents(context.Context, *ID) (*EventsResponse, error) {
	panic("GetUserEvents implement me")
}

// UpdateEvent updates event by id.
func (cs *CalendarServer) UpdateEvent(context.Context, *EventRequest) (*EventResponse, error) {
	panic("UpdateEvent implement me")
}

// DeleteEvent deletes event from DB.
func (cs *CalendarServer) DeleteEvent(context.Context, *ID) (*EventResponse, error) {
	panic("DeleteEvent implement me")
}

// GetEventsForDay returns all events for given day.
func (cs *CalendarServer) GetEventsForDay(context.Context, *Day) (*EventsResponse, error) {
	panic("GetEventsForDay implement me")
}

// GetEventsForWeek returns all events for given week from day.
func (cs *CalendarServer) GetEventsForWeek(context.Context, *Day) (*EventsResponse, error) {
	panic("GetEventsForWeek implement me")
}

// GetEventsForMonth returns all events for given month from day.
func (cs *CalendarServer) GetEventsForMonth(context.Context, *Day) (*EventsResponse, error) {
	panic("GetEventsForMonth implement me")
}
