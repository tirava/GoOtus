/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"github.com/golang/protobuf/ptypes"
	"sync"
)

// interface for any storage

var globID uint32

type dbEvents struct {
	sync.RWMutex
	events map[uint32]*Event
}

func newDBEvents() *dbEvents {
	return &dbEvents{
		events: make(map[uint32]*Event),
	}
}

func (db *dbEvents) addEvent(event *Event) {
	db.Lock()
	defer db.Unlock()
	db.events[event.Id] = event
}

func (db *dbEvents) delEvent(event *Event) {
	db.Lock()
	defer db.Unlock()
	db.events[event.Id].DeletedAt = ptypes.TimestampNow()
}

func (db *dbEvents) editEvent(event *Event) {
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = ptypes.TimestampNow()
	db.events[event.Id] = event
}

func newEvent() *Event {
	globID++
	return &Event{
		Id:        globID,
		CreatedAt: ptypes.TimestampNow(),
		UpdatedAt: ptypes.TimestampNow(),
		DeletedAt: nil,
		OccursAt:  nil,
		Subject:   "111",
		Body:      "222",
		Duration:  333,
		Location:  "Moscow",
		User: &User{
			Id:       1,
			Name:     "qqq",
			Email:    []string{"www"},
			Mobile:   []string{"+777"},
			Birthday: ptypes.TimestampNow(),
		},
	}
}
