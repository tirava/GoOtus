/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 29.10.2019 15:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package dbs implements db interfaces.
package dbs

import (
	"github.com/evakom/calendar/internal/dbs/inmemory"
	"github.com/evakom/calendar/internal/dbs/postgres"
	"github.com/evakom/calendar/internal/domain/interfaces/storage"
)

// Constants
const (
	MapDBType      = "map"
	PostgresDBType = "postgres"
	EventIDField   = "event_id"
	UserIDField    = "user_id"
	DayField       = "day"
)

// NewDB returns DB by db type
func NewDB(dbType string) (storage.DB, error) {
	switch dbType {
	case MapDBType:
		return inmemory.NewMapDB()
	case PostgresDBType:
		return postgres.NewPostgresDB()
	}
	return nil, nil
}
