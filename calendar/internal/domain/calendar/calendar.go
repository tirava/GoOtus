/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package calendar implements simple event calendar via protobuf.
package calendar

import (
	"context"
	"github.com/evakom/calendar/internal/domain/errors"
	"github.com/evakom/calendar/internal/domain/interfaces/storage"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
	"time"
)

// Calendar is the main calendar struct.
type Calendar struct {
	db storage.DB
}

// NewCalendar inits main calendar fields.
func NewCalendar(db storage.DB) Calendar {
	return Calendar{
		db: db,
	}
}

// AddEvent adds new event.
func (c Calendar) AddEvent(ctx context.Context, event models.Event) error {
	return c.db.AddEventDB(ctx, event)
}

// GetEvent got one event.
func (c Calendar) GetEvent(ctx context.Context, eventID uuid.UUID) (models.Event, error) {
	return c.db.GetOneEventDB(ctx, eventID)
}

// DelEvent deletes event.
func (c Calendar) DelEvent(ctx context.Context, eventID uuid.UUID) error {
	return c.db.DelEventDB(ctx, eventID)
}

// UpdateEvent updates event.
func (c Calendar) UpdateEvent(ctx context.Context, event models.Event) error {
	return c.db.EditEventDB(ctx, event)
}

// GetAllEventsFilter returns all calendar events with given filter.
func (c Calendar) GetAllEventsFilter(ctx context.Context, filter models.Event) ([]models.Event, error) {
	result := make([]models.Event, 0)
	dateNil := time.Time{}

	switch {
	case filter.ID != uuid.Nil:
		e, err := c.db.GetOneEventDB(ctx, filter.ID)
		if err != nil {
			return result, errors.ErrEventNotFound
		}
		result = append(result, e)
		return result, nil
	case filter.UserID != uuid.Nil:
		events := c.db.GetAllEventsDB(ctx, filter.UserID)
		if len(events) == 0 {
			return nil, errors.ErrEventsNotFound
		}
		return events, nil
	case filter.OccursAt != dateNil:
		events := c.db.GetAllEventsDBDays(ctx, filter)
		if len(events) == 0 {
			return nil, errors.ErrEventsNotFound
		}
		return events, nil
	default:
		return nil, errors.ErrNothingFound
	}
}

func (c Calendar) getEventUpdateTime(ctx context.Context, id uuid.UUID) (time.Time, error) {
	event, err := c.GetEvent(ctx, id)
	if err != nil {
		return time.Now(), errors.ErrEventNotFound
	}
	return event.UpdatedAt, nil
}

// UpdateEventFromEvent updates current event
// with fields from new event by event id
func (c Calendar) UpdateEventFromEvent(ctx context.Context, event models.Event) (models.Event, error) {

	e, err := c.GetEvent(ctx, event.ID)
	if err != nil {
		return event, errors.ErrEventNotFound
	}

	if event.Subject != "" {
		e.Subject = event.Subject
	}
	if event.Body != "" {
		e.Body = event.Body
	}
	if event.Location != "" {
		e.Location = event.Location
	}
	if event.Duration != 0 {
		e.Duration = event.Duration
	}
	timeNil := time.Time{}
	if event.OccursAt != timeNil {
		e.OccursAt = event.OccursAt
	}

	if err := c.UpdateEvent(ctx, e); err != nil {
		return e, err
	}
	if e.UpdatedAt, err = c.getEventUpdateTime(ctx, e.ID); err != nil {
		return e, err
	}

	return e, nil
}
