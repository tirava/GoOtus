/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package storage implements DB interfaces.
package storage

import (
	"context"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
)

// DB is thw main interface for any DBs
type DB interface {
	AddEventDB(context.Context, models.Event) error
	EditEventDB(context.Context, models.Event) error
	DelEventDB(context.Context, uuid.UUID) error
	GetOneEventDB(context.Context, uuid.UUID) (models.Event, error)
	GetAllEventsDB(context.Context, uuid.UUID) []models.Event
	CleanEventsDB(context.Context, uuid.UUID) error
	GetAllEventsDBDays(context.Context, models.Event) []models.Event
	CloseDB() error
}
