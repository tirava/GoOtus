/*
 * HomeWork-10: Calendar extending HTTP methods
 * Created on 17.11.2019 19:26
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package urlform implements www-url-from encode/decode of the models entities .
package urlform

import (
	"fmt"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"time"
)

// Constants.
const (
	FormOccursAt = "occurs_at"
	FormSubject  = "subject"
	FormBody     = "body"
	FormLocation = "location"
	FormDuration = "duration"
	FormUserID   = "user_id"
	FormEventID  = "event_id"
	FormDay      = "day"
	FormWeek     = "week"
	FormMonth    = "month"
)

// Values is the base www-url-form values type.
type Values map[string]string

// DecodeID returns decoded string id to uuid.
func DecodeID(sid string) (uuid.UUID, error) {
	uid := tools.IDString2UUIDorNil(sid)
	if uid == uuid.Nil {
		return uid, fmt.Errorf("invalid id=%s", sid)
	}
	return uid, nil
}

// DecodeEvent returns decoded event from www-url-form values.
func (v Values) DecodeEvent() (models.Event, error) {
	event := models.NewEvent()

	occurs, err := time.Parse("2006-01-02 15:04:05", v[FormOccursAt])
	if err != nil && v[FormOccursAt] != "" {
		return event, err
	}

	duration, err := time.ParseDuration(v[FormDuration])
	if err != nil && v[FormDuration] != "" {
		return event, err
	}

	if v[FormEventID] == "" {
		userID, err := DecodeID(v[FormUserID])
		if err != nil {
			return event, fmt.Errorf("illegal user id - %w", err)
		}
		event.UserID = userID
	} else {
		eventID, err := DecodeID(v[FormEventID])
		if err != nil {
			return event, fmt.Errorf("illegal event id - %w", err)
		}
		event.ID = eventID
	}

	event.Subject = v[FormSubject]
	event.Body = v[FormBody]
	event.Location = v[FormLocation]
	event.Duration = duration
	event.OccursAt = occurs

	return event, nil
}
