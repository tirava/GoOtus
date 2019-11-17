/*
 * HomeWork-10: Calendar extending HTTP methods
 * Created on 17.11.2019 19:26
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package urlform implements www-url-from encode/decode of the models entities .
package urlform

import (
	"errors"
	"github.com/evakom/calendar/internal/domain/models"
	"github.com/evakom/calendar/tools"
	"github.com/google/uuid"
	"time"
)

// Constants.
const (
	FormSubject  = "subject"
	FormBody     = "body"
	FormLocation = "location"
	FormDuration = "duration"
	FormUserID   = "user"
)

// Values is the base www-url-form values type.
type Values map[string]string

// DecodeEvent returns decoded event from www-url-form values.
func (v Values) DecodeEvent() (models.Event, error) {

	duration, err := time.ParseDuration(v[FormDuration])
	if err != nil {
		return models.Event{}, err
	}

	userID := tools.IDString2UUIDorNil(v[FormUserID])
	if userID == uuid.Nil {
		return models.Event{}, errors.New("illegal user id")
	}

	event := models.NewEvent()

	event.Subject = v[FormSubject]
	event.Body = v[FormBody]
	event.Location = v[FormLocation]
	event.Duration = duration
	event.UserID = userID

	return event, nil
}
