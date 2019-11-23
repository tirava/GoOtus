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
	AddEventDB(models.Event) error
	EditEventDB(models.Event) error
	DelEventDB(uuid.UUID) error
	GetOneEventDB(uuid.UUID) (models.Event, error)
	GetAllEventsDB(uuid.UUID) []models.Event
	CleanEventsDB(uuid.UUID) error
	GetAllEventsDBDays(models.Event) []models.Event
}
