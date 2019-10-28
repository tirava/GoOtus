/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 23.10.2019 19:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package interfaces implements interfaces.
package interfaces

import "github.com/evakom/calendar/internal/domain/models"

// Constants
const (
	MapDBType = "map"
	//PostgresDBType = "postgres"
)

// DB is thw main interface for any DBs
type DB interface {
	AddEvent(event models.Event) error
	EditEvent(event models.Event) error
	DelEvent(id int) error
	GetOneEvent(id int) (models.Event, error)
	GetAllEvents() []models.Event
}

// NewDB returns DB by db type
func NewDB(dbType string) DB {
	switch dbType {
	case MapDBType:
		return models.NewMapDB()
		//case PostgresDBType:
		//return models.NewPostgresDB()
	}
	return nil
}
