/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:12
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package inmemory implements memory DB interface.
package inmemory

import (
	"github.com/evakom/calendar/internal/domain/errors"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/internal/loggers"
	"github.com/google/uuid"
	"sync"
	"time"
)

// Constants
const (
	EventIDField = "event_id"
	UserIDField  = "user_id"
	DayField     = "day"
	DeltaField   = "delta"
)

// DBMapEvents is the base struct for using map db.
type DBMapEvents struct {
	sync.RWMutex
	events map[uuid.UUID]models.Event
	logger loggers.Logger
}

// NewMapDB returns new map db struct.
func NewMapDB() (*DBMapEvents, error) {
	dbm := &DBMapEvents{
		events: make(map[uuid.UUID]models.Event),
		logger: loggers.GetLogger(),
	}
	dbm.logger.Info("New map DB created")
	return dbm, nil
}

// AddEventDB adds event to map db.
func (db *DBMapEvents) AddEventDB(event models.Event) error {
	db.Lock()
	defer db.Unlock()
	if _, ok := db.events[event.ID]; ok {
		return errors.ErrEventAlreadyExists
	}
	db.events[event.ID] = event
	db.logger.WithFields(loggers.Fields{
		EventIDField: event.ID.String(),
		UserIDField:  event.UserID.String(),
	}).Info("Event added into map DB")
	db.logger.Debug("Event body added into map DB: %+v", event)
	return nil
}

// DelEventDB deletes one event by id.
func (db *DBMapEvents) DelEventDB(id uuid.UUID) error {
	if _, ok := db.events[id]; !ok {
		return errors.ErrEventNotFound
	}
	db.Lock()
	defer db.Unlock()
	e := db.events[id]
	e.DeletedAt = time.Now()
	db.events[id] = e
	db.logger.WithFields(loggers.Fields{
		EventIDField: id.String(),
		UserIDField:  e.UserID.String(),
	}).Info("Event deleted from map DB")
	db.logger.Debug("Event body deleted from map DB: %+v", e)
	return nil
}

// EditEventDB updates one event.
func (db *DBMapEvents) EditEventDB(event models.Event) error {
	if _, ok := db.events[event.ID]; !ok {
		return errors.ErrEventNotFound
	}
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = time.Now()
	db.events[event.ID] = event
	db.logger.WithFields(loggers.Fields{
		EventIDField: event.ID.String(),
		UserIDField:  event.UserID.String(),
	}).Info("Event updated in map DB")
	db.logger.Debug("Event body updated in map DB: %+v", event)
	return nil
}

// GetOneEventDB returns one event by id.
func (db *DBMapEvents) GetOneEventDB(id uuid.UUID) (models.Event, error) {
	if _, ok := db.events[id]; !ok {
		return models.Event{}, errors.ErrEventNotFound
	}
	if !db.events[id].DeletedAt.IsZero() {
		return models.Event{}, errors.ErrEventAlreadyDeleted
	}
	db.logger.WithFields(loggers.Fields{
		EventIDField: id.String(),
		UserIDField:  db.events[id].UserID.String(),
	}).Info("Event got from map DB")
	db.logger.Debug("Event body got from map DB: %+v", db.events[id])
	return db.events[id], nil
}

// GetAllEventsDB return all events slice for given user id (no deleted).
func (db *DBMapEvents) GetAllEventsDB(id uuid.UUID) []models.Event {
	events := make([]models.Event, 0, len(db.events))
	for _, event := range db.events {
		if !event.DeletedAt.IsZero() {
			continue
		}
		if event.UserID == id || id == uuid.Nil {
			events = append(events, event)
		}
	}
	db.logger.WithFields(loggers.Fields{
		UserIDField: id.String(),
	}).Info("All events [%d] got from map DB", len(events))
	return events
}

// CleanEventsDB cleans db and deletes all events in the db for given user id (no restoring!).
func (db *DBMapEvents) CleanEventsDB(id uuid.UUID) error {
	db.Lock()
	defer db.Unlock()
	if id == uuid.Nil {
		db.events = make(map[uuid.UUID]models.Event)
	} else {
		for _, event := range db.events {
			if event.UserID == id {
				delete(db.events, event.ID)
			}
		}
	}
	db.logger.WithFields(loggers.Fields{
		UserIDField: id.String(),
	}).Info("All events deleted in map DB")
	return nil
}

// GetAllEventsDBDays returns events for num of the days for given user
func (db *DBMapEvents) GetAllEventsDBDays(filter models.Event) []models.Event {
	events := make([]models.Event, 0, len(db.events))
	occursEnd := filter.OccursAt.Add(filter.Duration)
	for _, event := range db.events {
		if event.UserID != filter.UserID && filter.UserID != uuid.Nil {
			continue
		}
		if event.OccursAt.Equal(filter.OccursAt) ||
			(event.OccursAt.After(filter.OccursAt) && event.OccursAt.Before(occursEnd)) {
			events = append(events, event)
		}
	}
	db.logger.WithFields(loggers.Fields{
		DayField:   filter.OccursAt.String(),
		DeltaField: filter.Duration,
	}).Info("All events [%d] for day(s) with delta got from map DB", len(events))
	return events
}
