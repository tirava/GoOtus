/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 29.10.2019 16:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package dbs

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	_ "github.com/jackc/pgx/stdlib" // driver for postgres
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

// TODO into config
const dsn = ""

// DBPostgresEvents is the base struct for using map db.
type DBPostgresEvents struct {
	db *sqlx.DB
}

// NewPostgresDB returns new postgres db struct.
func NewPostgresDB() (*DBPostgresEvents, error) {
	db, err := sqlx.Open("pgx", dsn)
	if err != nil {
		return &DBPostgresEvents{}, fmt.Errorf("error open db: %w", err)
	}
	err = db.Ping()
	if err != nil {
		return &DBPostgresEvents{}, fmt.Errorf("error ping db: %w", err)
	}
	return &DBPostgresEvents{db: db}, nil
}

// AddEvent adds event to postgres db.
func (db *DBPostgresEvents) AddEvent(event models.Event) error {
	// TODO
	return nil
}

// DelEvent deletes one event by id.
func (db *DBPostgresEvents) DelEvent(id uuid.UUID) error {
	// TODO
	return nil
}

// EditEvent updates one event.
func (db *DBPostgresEvents) EditEvent(event models.Event) error {
	// TODO
	return nil
}

// GetOneEvent returns one event by id.
func (db *DBPostgresEvents) GetOneEvent(id uuid.UUID) (models.Event, error) {
	// TODO
	return models.Event{}, nil
}

// GetAllEvents return all events slice.
func (db *DBPostgresEvents) GetAllEvents() []models.Event {
	// TODO
	return []models.Event{}
}
