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

var globID uint32

type db interface {
	newDB() interface{}
	addEvent(event *Event) error
	editEvent(event *Event) error
	delEvent(id uint32) error
}

func newDB(d db) interface{} {
	return d.newDB()
}

type dbMapEvents struct {
	sync.RWMutex
	events map[uint32]*Event
}

func (db *dbMapEvents) newDB() interface{} {
	return &dbMapEvents{
		events: make(map[uint32]*Event),
	}
}

func (db *dbMapEvents) addEvent(event *Event) error {
	db.Lock()
	defer db.Unlock()
	db.events[event.Id] = event
	return nil
}

func (db *dbMapEvents) delEvent(id uint32) error {
	db.Lock()
	defer db.Unlock()
	db.events[id].DeletedAt = ptypes.TimestampNow()
	return nil
}

func (db *dbMapEvents) editEvent(event *Event) error {
	db.Lock()
	defer db.Unlock()
	event.UpdatedAt = ptypes.TimestampNow()
	db.events[event.Id] = event
	return nil
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
