/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

package main

import (
	"github.com/golang/protobuf/ptypes"
)

// interface for any storage
// mutex for edit

//type dbEvents struct {
//	sync.Mutex
//	Events
//}

func (m *Events) addEvent(event *Event) {
	m.Events = append(m.Events, event)
}

func newEvent() *Event {
	return &Event{
		Id:        1,
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
