/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:11
 * Copyright (c) 2019 - Eugene Klimov
 */

package calendar

import "github.com/golang/protobuf/ptypes"

var globID uint32

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
