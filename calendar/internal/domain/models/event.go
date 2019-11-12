/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 24.10.2019 19:11
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package models implements base models.
package models

import (
	"fmt"
	"github.com/google/uuid"
	"strings"
	"time"
)

// FORMATSTRING for stringer
const FORMATSTRING = "%-13s%s\n"

// Event is the base event struct.
type Event struct {
	ID        uuid.UUID
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	OccursAt  time.Time
	Subject   string
	Body      string
	Duration  time.Duration
	Location  string
	UserID    uuid.UUID
}

// NewEvent returns new example event.
func NewEvent() Event {
	return Event{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Subject:   "111",
		Body:      "222",
		Duration:  time.Minute,
		Location:  "Moscow",
		UserID:    uuid.New(), // todo uuid.Nil
	}
}

// StringEr is event stringer
func (e Event) StringEr() string {
	sb := strings.Builder{}
	sb.WriteString(fmt.Sprintf(FORMATSTRING, "ID:", e.ID))
	sb.WriteString(fmt.Sprintf(FORMATSTRING, "CreatedAt:", e.CreatedAt))
	sb.WriteString(fmt.Sprintf(FORMATSTRING, "UpdatedAt:", e.UpdatedAt))
	sb.WriteString(fmt.Sprintf(FORMATSTRING, "Subject:", e.Subject))
	sb.WriteString(fmt.Sprintf(FORMATSTRING, "Body:", e.Body))
	sb.WriteString(fmt.Sprintf(FORMATSTRING, "Duration:", e.Duration))
	sb.WriteString(fmt.Sprintf(FORMATSTRING, "Location:", e.Location))
	sb.WriteString(fmt.Sprintf(FORMATSTRING, "UserID:", e.UserID))
	return sb.String()
}
