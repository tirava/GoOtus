/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 29.10.2019 15:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package dbs implements db interfaces.
package dbs

import (
	"context"
	"github.com/evakom/calendar/internal/dbs/inmemory"
	"github.com/evakom/calendar/internal/dbs/postgres"
	"github.com/evakom/calendar/internal/domain/interfaces/storage"
)

// Constants
const (
	MapDBType      = "map"
	PostgresDBType = "postgres"
)

// NewDB returns DB by db type
func NewDB(dbType, dsn string, ctx context.Context) (storage.DB, error) {
	switch dbType {
	case MapDBType:
		return inmemory.NewMapDB()
	case PostgresDBType:
		return postgres.NewPostgresDB(dsn, ctx)
	}
	return nil, nil
}
