/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package storage implements DB interfaces.
package storage

import (
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
)

// DB is thw main interface for any DBs
type DB interface {
	AddEventDB(event models.Event) error
	EditEventDB(event models.Event) error
	DelEventDB(id uuid.UUID) error
	GetOneEventDB(id uuid.UUID) (models.Event, error)
	GetAllEventsDB(id uuid.UUID) []models.Event
	CleanEventsDB(id uuid.UUID) error
}