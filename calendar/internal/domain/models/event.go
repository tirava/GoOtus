/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:11
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package models implements base models.
package models

import (
	"time"
)

var globID int

// Event is the base event struct.
type Event struct {
	ID        int
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	OccursAt  time.Time
	Subject   string
	Body      string
	Duration  int
	Location  string
	User      User
}

// NewEvent returns new example event.
func NewEvent() Event {
	globID++
	return Event{
		ID:        globID,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Subject:   "111",
		Body:      "222",
		Duration:  333,
		Location:  "Moscow",
		User: User{
			ID:       1,
			Name:     "qqq",
			Email:    []string{"www"},
			Mobile:   []string{"+777"},
			Birthday: time.Now(),
		},
	}
}
