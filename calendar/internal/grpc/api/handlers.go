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
	UserIDField  = "user_id"
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

	if protoEvent.CreatedAt, err = ptypes.TimestampProto(event.CreatedAt); err != nil {
		cs.logger.Error("[CreateEvent] error convert event time to proto: %s", err)
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
func (cs *CalendarServer) GetEvent(ctx context.Context, id *ID) (*EventResponse, error) {
	cs.logger.Info("REQUEST [GetEvent]")

	eid, err := uuid.Parse(id.GetId())
	if err != nil {
		cs.logger.WithFields(loggers.Fields{
			CodeField: codes.InvalidArgument,
		}).Error("RESPONSE [GetEvent]: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	event, err := cs.calendar.GetEvent(eid)
	if err != nil {
		cs.logger.WithFields(loggers.Fields{
			CodeField: codes.Internal,
		}).Error("RESPONSE [GetEvent]: %s", err)
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

	createdAt, err := ptypes.TimestampProto(event.CreatedAt)
	if err != nil {
		cs.logger.Error("[GetEvent] error convert event create to proto: %s", err)
	}
	updatedAt, err := ptypes.TimestampProto(event.UpdatedAt)
	if err != nil {
		cs.logger.Error("[GetEvent] error convert event update to proto: %s", err)
	}
	occursAt, err := ptypes.TimestampProto(event.OccursAt)
	if err != nil {
		cs.logger.Error("[GetEvent] error convert event occurs to proto: %s", err)
	}

	protoEvent := &Event{
		Id:        event.ID.String(),
		CreatedAt: createdAt,
		UpdatedAt: updatedAt,
		OccursAt:  occursAt,
		Subject:   event.Subject,
		Body:      event.Body,
		Duration:  ptypes.DurationProto(event.Duration),
		Location:  event.Location,
		UserID:    event.UserID.String(),
	}

	cs.logger.WithFields(loggers.Fields{
		CodeField:    codes.OK,
		EventIDField: protoEvent.Id,
	}).Info("RESPONSE [GetEvent]")

	resp := &EventResponse{
		Result: &EventResponse_Event{
			Event: protoEvent,
		},
	}
	cs.logger.Debug("[GetEvent] Response body: %+v", resp)

	return resp, nil
}

// GetUserEvents returns all events for given user.
func (cs *CalendarServer) GetUserEvents(ctx context.Context, id *ID) (*EventsResponse, error) {
	cs.logger.Info("REQUEST [GetUserEvents]")

	uid, err := uuid.Parse(id.GetId())
	if err != nil {
		cs.logger.WithFields(loggers.Fields{
			CodeField: codes.InvalidArgument,
		}).Error("RESPONSE [GetUserEvents]: %s", err)
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	events, err := cs.calendar.GetAllEventsFilter(models.Event{UserID: uid})
	if err != nil {
		cs.logger.WithFields(loggers.Fields{
			CodeField: codes.Internal,
		}).Error("RESPONSE [GetUserEvents]: %s", err)
		return nil, status.Error(codes.Internal, err.Error())
	}

	protoEvents := make([]*Event, 0)

	for _, event := range events {
		createdAt, err := ptypes.TimestampProto(event.CreatedAt)
		if err != nil {
			cs.logger.Error("[GetUserEvents] error convert event create to proto: %s", err)
		}
		updatedAt, err := ptypes.TimestampProto(event.UpdatedAt)
		if err != nil {
			cs.logger.Error("[GetUserEvents] error convert event update to proto: %s", err)
		}
		occursAt, err := ptypes.TimestampProto(event.OccursAt)
		if err != nil {
			cs.logger.Error("[GetUserEvents] error convert event occurs to proto: %s", err)
		}
		protoEvent := &Event{
			Id:        event.ID.String(),
			CreatedAt: createdAt,
			UpdatedAt: updatedAt,
			OccursAt:  occursAt,
			Subject:   event.Subject,
			Body:      event.Body,
			Duration:  ptypes.DurationProto(event.Duration),
			Location:  event.Location,
			UserID:    event.UserID.String(),
		}
		protoEvents = append(protoEvents, protoEvent)
	}

	cs.logger.WithFields(loggers.Fields{
		CodeField:   codes.OK,
		UserIDField: id.GetId(),
	}).Info("RESPONSE [GetUserEvents]")

	return &EventsResponse{Events: protoEvents}, nil
}

// DeleteEvent deletes event from DB.
func (cs *CalendarServer) DeleteEvent(context.Context, *ID) (*EventResponse, error) {
	panic("DeleteEvent implement me")
}

// UpdateEvent updates event by id.
func (cs *CalendarServer) UpdateEvent(context.Context, *EventRequest) (*EventResponse, error) {
	panic("UpdateEvent implement me")
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
