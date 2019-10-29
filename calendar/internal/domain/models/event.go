/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:11
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package models implements base models.
package models

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

// Event is the base event struct.
type Event struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	OccursAt  time.Time
	Subject   string
	Body      string
	Duration  time.Duration
	Location  string
}

// NewEvent returns new example event.
func NewEvent() Event {
	return Event{
		ID:        uuid.NewV4(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Subject:   "111",
		Body:      "222",
		Duration:  time.Minute,
		Location:  "Moscow",
	}
}
