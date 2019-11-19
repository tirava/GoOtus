/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package calendar implements simple event calendar via protobuf.
package calendar

import (
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
func (c Calendar) AddEvent(event models.Event) error {
	return c.db.AddEventDB(event)
}

// GetEvent got one event.
func (c Calendar) GetEvent(eventID uuid.UUID) (models.Event, error) {
	return c.db.GetOneEventDB(eventID)
}

// DelEvent deletes event.
func (c Calendar) DelEvent(eventID uuid.UUID) error {
	return c.db.DelEventDB(eventID)
}

// UpdateEvent updates event.
func (c Calendar) UpdateEvent(event models.Event) error {
	return c.db.EditEventDB(event)
}

// GetAllEventsFilter returns all calendar events with given filter.
func (c Calendar) GetAllEventsFilter(filter models.Event, opt int) ([]models.Event, error) {
	result := make([]models.Event, 0)

	if filter.ID != uuid.Nil {
		e, err := c.db.GetOneEventDB(filter.ID)
		if err != nil {
			return result, errors.ErrEventNotFound
		}
		result = append(result, e)
		return result, nil
	}

	if filter.UserID != uuid.Nil {
		events := c.db.GetAllEventsDB(filter.UserID)
		if len(events) == 0 {
			return nil, errors.ErrEventsNotFound
		}
		return events, nil
	}

	dNil := time.Time{}
	if filter.OccursAt != dNil {
		events := c.db.GetAllEventsDBDays(filter.OccursAt, opt)
		if len(events) == 0 {
			return nil, errors.ErrEventsNotFound
		}
		return events, nil
	}

	return nil, errors.ErrNothingFound
}

func (c Calendar) getEventUpdateTime(id uuid.UUID) (time.Time, error) {
	event, err := c.GetEvent(id)
	if err != nil {
		return time.Now(), errors.ErrEventNotFound
	}
	return event.UpdatedAt, nil
}

// UpdateEventFromEvent updates current event
// with fields from new event by event id
func (c Calendar) UpdateEventFromEvent(event models.Event) (models.Event, error) {

	e, err := c.GetEvent(event.ID)
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

	if err := c.UpdateEvent(e); err != nil {
		return e, err
	}
	if e.UpdatedAt, err = c.getEventUpdateTime(e.ID); err != nil {
		return e, err
	}

	return e, nil
}
