/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package dbs

import (
	"github.com/evakom/calendar/internal/domain/errors"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
	"sync"
	"time"
)

// DBMapEvents is the base struct for using map db.
type DBMapEvents struct {
	sync.RWMutex
	events map[uuid.UUID]models.Event
	logger models.Logger
}

// NewMapDB returns new map db struct.
func NewMapDB() (*DBMapEvents, error) {
	dbm := &DBMapEvents{
		events: make(map[uuid.UUID]models.Event),
		logger: models.Logger{}.GetLogger(),
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
	db.logger.WithFields(models.Fields{
		"id": event.ID.String(),
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
	db.logger.WithFields(models.Fields{
		"id": id.String(),
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
	db.logger.WithFields(models.Fields{
		"id": event.ID.String(),
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
	db.logger.WithFields(models.Fields{
		"id": id.String(),
	}).Info("Event got from map DB")
	db.logger.Debug("Event body got from map DB: %+v", db.events[id])
	return db.events[id], nil
}

// GetAllEventsDB return all events slice (no deleted).
func (db *DBMapEvents) GetAllEventsDB() []models.Event {
	events := make([]models.Event, 0, len(db.events))
	for _, event := range db.events {
		if !event.DeletedAt.IsZero() {
			continue
		}
		events = append(events, event)
	}
	db.logger.Info("All events got from map DB")
	return events
}

// CleanEventsDB cleans db and deletes all events in the db (no restoring!).
func (db *DBMapEvents) CleanEventsDB() error {
	db.Lock()
	defer db.Unlock()
	db.events = make(map[uuid.UUID]models.Event)
	db.logger.Info("Map DB cleaned, all events deleted")
	return nil
}
