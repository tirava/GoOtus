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

// AddEvent adds new event for given user.
func (c Calendar) AddEvent(event models.Event) error {
	return c.db.AddEventDB(event)
}

// GetAllEventsFilter returns all calendar events with given filter.
func (c Calendar) GetAllEventsFilter(filter models.Event) ([]models.Event, error) {
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

	return nil, nil
}
