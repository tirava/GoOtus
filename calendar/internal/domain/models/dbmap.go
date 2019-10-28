/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package models

import (
	"github.com/evakom/calendar/internal/domain/errors"
	"sync"
	"time"
)

// DBMapEvents is the base struct for using map db.
type DBMapEvents struct {
	sync.RWMutex
	Events map[int]Event
}

// NewMapDB returns new map db struct.
func NewMapDB() *DBMapEvents {
	return &DBMapEvents{
		Events: make(map[int]Event),
	}
}

// AddEvent adds event to map db.
func (db *DBMapEvents) AddEvent(event Event) error {
	db.Lock()
	defer db.Unlock()
	db.Events[event.ID] = event
	return nil
}

// DelEvent deletes one event by id.
func (db *DBMapEvents) DelEvent(id int) error {
	if _, ok := db.Events[id]; !ok {
		return errors.ErrEventNotFound
		//return fmt.Errorf("event id = %d not found", id)
	}
	db.Lock()
	defer db.Unlock()
	e := db.Events[id]
	e.DeletedAt = time.Now()
	db.Events[id] = e
	return nil
}

// EditEvent updates one event.
func (db *DBMapEvents) EditEvent(event Event) error {
	if _, ok := db.Events[event.ID]; !ok {
		return errors.ErrEventNotFound
		//return fmt.Errorf("event id = %d not found", event.ID)
	}
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = time.Now()
	db.Events[event.ID] = event
	return nil
}

// GetOneEvent returns one event by id.
func (db *DBMapEvents) GetOneEvent(id int) (Event, error) {
	if _, ok := db.Events[id]; !ok {
		return Event{}, errors.ErrEventNotFound
		//return Event{}, fmt.Errorf("event id = %d not found", id)
	}
	if !db.Events[id].DeletedAt.IsZero() {
		return Event{}, errors.ErrEventAlreadyDeleted
		//return Event{}, fmt.Errorf("event id = %d already deleted", id)
	}
	return db.Events[id], nil
}

// GetAllEvents return all events slice.
func (db *DBMapEvents) GetAllEvents() []Event {
	events := make([]Event, 0)
	for _, event := range db.Events {
		events = append(events, event)
	}
	return events
}
