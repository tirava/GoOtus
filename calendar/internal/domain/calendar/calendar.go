/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 22.10.2019 22:44
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package calendar implements simple event calendar via protobuf.
package calendar

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/interfaces"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/google/uuid"
)

// Calendar is the main calendar struct.
type Calendar struct {
	db     interfaces.DB
	logger models.Logger
}

// NewCalendar inits main calendar fields.
func NewCalendar(db interfaces.DB) Calendar {
	return Calendar{
		db:     db,
		logger: models.Logger{}.GetLogger(),
	}
}

// GetAllEvents returns all calendar events for given user
func (c Calendar) GetAllEvents(userID uuid.UUID) string {

	event1 := models.NewEvent()
	event1.Location = "qqqqqqqqqqqqqqqqqqqqqq"
	_ = c.db.AddEvent(event1)

	event2 := models.NewEvent()
	event2.Subject = "222222222222222222222"
	event2.Body = "3333333333333333333"
	_ = c.db.AddEvent(event2)

	return fmt.Sprint(c.db.GetAllEvents())
}

// PrintTestData test events.
//func (c Calendar) PrintTestData() {
//	event1 := models.NewEvent()
//
//	event1.Location = "qqqqqqqqqqqqqqqqqqqqqq"
//	_ = c.db.AddEvent(event1)
//
//	event2 := models.NewEvent()
//	event2.Subject = "222222222222222222222"
//	event2.Body = "3333333333333333333"
//	_ = c.db.AddEvent(event2)
//
//	c.logger.Info("%+v\n", c.db.GetAllEvents())
//
//	if err := c.db.DelEvent(event1.ID); err != nil {
//		c.logger.Error(err.Error())
//	}
//
//	c.logger.Info("%+v\n", c.db.GetAllEvents())
//
//	event2.Duration = time.Hour * 8
//	if err := c.db.EditEvent(event2); err != nil {
//		c.logger.Error(err.Error())
//	}
//
//	c.logger.Info("%+v\n", c.db.GetAllEvents())
//
//	e2, err := c.db.GetOneEvent(event2.ID)
//	if err != nil {
//		c.logger.Error(err.Error())
//	}
//
//	c.logger.Info("%+v\n", e2)
//}
