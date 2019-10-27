/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

import (
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"sync"
)

// DBMapEvents is the base struct for using map db.
type DBMapEvents struct {
	sync.RWMutex
	Events map[uint32]Event
}

func newMapDB() *DBMapEvents {
	return &DBMapEvents{
		Events: make(map[uint32]Event),
	}
}

// AddEvent adds event to map db.
func (db *DBMapEvents) AddEvent(event Event) error {
	db.Lock()
	defer db.Unlock()
	db.Events[event.Id] = event
	return nil
}

// DelEvent deletes one event by id
func (db *DBMapEvents) DelEvent(id uint32) error {
	if _, ok := db.Events[id]; !ok {
		return fmt.Errorf("event id = %d not found", id)
	}
	db.Lock()
	defer db.Unlock()
	e := db.Events[id]
	e.DeletedAt = ptypes.TimestampNow()
	db.Events[id] = e
	//db.Events[id].DeletedAt = ptypes.TimestampNow()
	return nil
}

// EditEvent updates one event.
func (db *DBMapEvents) EditEvent(event Event) error {
	if _, ok := db.Events[event.Id]; !ok {
		return fmt.Errorf("event id = %d not found", event.Id)
	}
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = ptypes.TimestampNow()
	db.Events[event.Id] = event
	return nil
}

// GetEvent returns one event by id.
func (db *DBMapEvents) GetEvent(id uint32) (Event, error) {
	if _, ok := db.Events[id]; !ok {
		return Event{}, fmt.Errorf("event id = %d not found", id)
	}
	if db.Events[id].DeletedAt != nil {
		return Event{}, fmt.Errorf("event id = %d already deleted", id)
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
