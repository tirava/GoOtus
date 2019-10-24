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

type dbMapEvents struct {
	sync.RWMutex
	events map[uint32]*Event
}

func (db *dbMapEvents) newDB() interface{} {
	return &dbMapEvents{
		events: make(map[uint32]*Event),
	}
}

func (db *dbMapEvents) addEvent(event Event) error {
	db.Lock()
	defer db.Unlock()
	db.events[event.Id] = &event
	return nil
}

func (db *dbMapEvents) delEvent(id uint32) error {
	if _, ok := db.events[id]; !ok {
		return fmt.Errorf("event id = %d not found", id)
	}
	db.Lock()
	defer db.Unlock()
	db.events[id].DeletedAt = ptypes.TimestampNow()
	return nil
}

func (db *dbMapEvents) editEvent(event Event) error {
	if _, ok := db.events[event.Id]; !ok {
		return fmt.Errorf("event id = %d not found", event.Id)
	}
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = ptypes.TimestampNow()
	db.events[event.Id] = &event
	return nil
}

func (db *dbMapEvents) getEvent(id uint32) (Event, error) {
	if _, ok := db.events[id]; !ok {
		return Event{}, fmt.Errorf("event id = %d not found", id)
	}
	if db.events[id].DeletedAt != nil {
		return Event{}, fmt.Errorf("event id = %d deleted", id)
	}
	return *db.events[id], nil
}

func (db *dbMapEvents) getAllEvents() []Event {
	events := make([]Event, 0)
	for _, event := range db.events {
		events = append(events, *event)
	}
	return events
}
