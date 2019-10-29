/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package dbs

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	uuid "github.com/satori/go.uuid"
	"sync"
	"time"
)

// DBMapEvents is the base struct for using map db.
type DBMapEvents struct {
	sync.RWMutex
	events map[uuid.UUID]models.Event
}

// NewMapDB returns new map db struct.
func NewMapDB() (*DBMapEvents, error) {
	return &DBMapEvents{
		events: make(map[uuid.UUID]models.Event),
	}, nil
}

// AddEvent adds event to map db.
func (db *DBMapEvents) AddEvent(event models.Event) error {
	db.Lock()
	defer db.Unlock()
	if _, ok := db.events[event.ID]; ok {
		return fmt.Errorf("event id = %s already exists", event.ID.String())
	}
	db.events[event.ID] = event
	return nil
}

// DelEvent deletes one event by id.
func (db *DBMapEvents) DelEvent(id uuid.UUID) error {
	if _, ok := db.events[id]; !ok {
		return fmt.Errorf("event id = %s not found", id.String())
	}
	db.Lock()
	defer db.Unlock()
	e := db.events[id]
	e.DeletedAt = time.Now()
	db.events[id] = e
	return nil
}

// EditEvent updates one event.
func (db *DBMapEvents) EditEvent(event models.Event) error {
	if _, ok := db.events[event.ID]; !ok {
		return fmt.Errorf("event id = %s not found", event.ID.String())
	}
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = time.Now()
	db.events[event.ID] = event
	return nil
}

// GetOneEvent returns one event by id.
func (db *DBMapEvents) GetOneEvent(id uuid.UUID) (models.Event, error) {
	if _, ok := db.events[id]; !ok {
		return models.Event{}, fmt.Errorf("event id = %d not found", id)
	}
	if !db.events[id].DeletedAt.IsZero() {
		return models.Event{}, fmt.Errorf("event id = %d already deleted", id)
	}
	return db.events[id], nil
}

// GetAllEvents return all events slice (no deleted).
func (db *DBMapEvents) GetAllEvents() []models.Event {
	events := make([]models.Event, 0, len(db.events))
	for _, event := range db.events {
		if !event.DeletedAt.IsZero() {
			continue
		}
		events = append(events, event)
	}
	return events
}
