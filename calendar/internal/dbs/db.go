/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 29.10.2019 15:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package dbs implements db interfaces.
package dbs

import (
	"github.com/evakom/calendar/internal/domain/interfaces"
)

// Constants
const (
	MapDBType      = "map"
	PostgresDBType = "postgres"
)

// NewDB returns DB by db type
func NewDB(dbType string) (interfaces.DB, error) {
	switch dbType {
	case MapDBType:
		return NewMapDB()
	case PostgresDBType:
		return NewPostgresDB()
	}
	return nil, nil
}
