/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package calendar implements simple event calendar via protobuf.
package calendar

import (
	"github.com/evakom/calendar/internal/domain/interfaces"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
)

// Calendar is the main calendar struct.
type Calendar struct {
	db     interfaces.DB
	logger models.Logger
}

// NewCalendar inits main calendar fields.
func NewCalendar(db interfaces.DB) Calendar {
	return Calendar{
		db:     db,
		logger: models.Logger{}.GetLogger(),
	}
}

// AddEvent adds new event for given user
func (c Calendar) AddEvent(event models.Event) error {
	return c.db.AddEventDB(event)
}

// GetAllEvents returns all calendar events for given user
func (c Calendar) GetAllEvents(userID string) []models.Event {
	uid, err := uuid.Parse(userID)
	if err != nil {
		return nil
	}

	return c.db.GetAllEventsDB()
}
