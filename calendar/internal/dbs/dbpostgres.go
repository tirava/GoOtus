/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 29.10.2019 16:12
 * Copyright (c) 2019 - Eugene Klimov
 */

package dbs

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
	_ "github.com/jackc/pgx/stdlib" // driver for postgres
	"github.com/jmoiron/sqlx"
)

// TODO into config
const dsn = ""

// DBPostgresEvents is the base struct for using map db.
type DBPostgresEvents struct {
	db     *sqlx.DB
	logger models.Logger
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
	dbp := &DBPostgresEvents{
		db:     db,
		logger: models.Logger{}.GetLogger(),
	}
	dbp.logger.Info("Connected to postgres DB")
	return dbp, nil
}

// AddEventDB adds event to postgres db.
func (db *DBPostgresEvents) AddEventDB(event models.Event) error {
	// TODO
	db.logger.WithFields(models.Fields{
		"id": event.ID.String(),
	}).Info("Event added into postgres DB")
	db.logger.Debug("Event body added: %+v", event)
	return nil
}

// DelEventDB deletes one event by id.
func (db *DBPostgresEvents) DelEventDB(id uuid.UUID) error {
	// TODO
	db.logger.WithFields(models.Fields{
		"id": id.String(),
	}).Info("Event deleted from postgres DB")
	//db.logger.Debug("Event body deleted from postgres DB: %+v", e)
	return nil
}

// EditEventDB updates one event.
func (db *DBPostgresEvents) EditEventDB(event models.Event) error {
	// TODO
	db.logger.WithFields(models.Fields{
		"id": event.ID.String(),
	}).Info("Event updated in postgres DB")
	db.logger.Debug("Event body updated in postgres DB: %+v", event)
	return nil
}

// GetOneEventDB returns one event by id.
func (db *DBPostgresEvents) GetOneEventDB(id uuid.UUID) (models.Event, error) {
	// TODO
	db.logger.WithFields(models.Fields{
		"id": id.String(),
	}).Info("Event got from postgres DB")
	//db.logger.Debug("Event body got from postgres DB: %+v", db.events[id])
	return models.Event{}, nil
}

// GetAllEventsDB return all events slice.
func (db *DBPostgresEvents) GetAllEventsDB() []models.Event {
	// TODO
	db.logger.Info("All events got from postgres DB")
	return []models.Event{}
}

// CleanEventsDB cleans db and deletes all events in the db (no restoring!).
func (db *DBPostgresEvents) CleanEventsDB() error {
	// TODO
	db.logger.Info("Postgres DB cleaned, all events deleted")
	return nil
}
