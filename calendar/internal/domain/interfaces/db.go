/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package interfaces implements interfaces.
package interfaces

import (
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
)

// DB is thw main interface for any DBs
type DB interface {
	AddEvent(event models.Event) error
	EditEvent(event models.Event) error
	DelEvent(id uuid.UUID) error
	GetOneEvent(id uuid.UUID) (models.Event, error)
	GetAllEvents() []models.Event
}
