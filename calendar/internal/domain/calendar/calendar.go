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
	"github.com/evakom/calendar/internal/loggers"
	"github.com/google/uuid"
)

// Calendar is the main calendar struct.
type Calendar struct {
	db     storage.DB
	logger loggers.Logger
}

// NewCalendar inits main calendar fields.
func NewCalendar(db storage.DB) Calendar {
	return Calendar{
		db:     db,
		logger: loggers.Logger{}.GetLogger(),
	}
}

// AddEvent adds new event for given user.
func (c Calendar) AddEvent(event models.Event) error {
	c.logger.WithFields(loggers.Fields{
		"id":     event.ID.String(),
		"userID": event.UserID.String(),
	}).Info("Request add event into calendar")
	c.logger.Debug("Requested event body for adding into calendar: %+v", event)
	return c.db.AddEventDB(event)
}

// GetAllEventsFilter returns all calendar events with given filter.
func (c Calendar) GetAllEventsFilter(filter models.Event) ([]models.Event, error) {
	result := make([]models.Event, 0)

	if filter.ID != uuid.Nil {
		e, err := c.db.GetOneEventDB(filter.ID)
		if err != nil {
			c.logger.WithFields(loggers.Fields{
				"id": filter.ID,
			}).Error("Filtered error: %s", err.Error())
			return result, errors.ErrEventNotFound
		}
		result = append(result, e)
		c.logger.WithFields(loggers.Fields{
			"id": filter.ID,
		}).Info("Returned filtered event by eventID")
		return result, nil
	}

	if filter.UserID != uuid.Nil {
		events := c.db.GetAllEventsDB()
		for _, e := range events {
			if e.UserID == filter.UserID {
				result = append(result, e)
			}
		}
		c.logger.WithFields(loggers.Fields{
			"userID": filter.UserID,
		}).Info("Returned filtered events by userID")
		return result, nil
	}

	c.logger.Warn("No returned events by filter")
	return nil, nil
}
