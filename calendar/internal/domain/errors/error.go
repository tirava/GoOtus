/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 28.10.2019 22:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package errors implements base calendar errors.
package errors

// EventError is the base type foe events errors.
type EventError string

// Error returns string error of the EventError.
func (e EventError) Error() string {
	return string(e)
}

// Errors
var (
	ErrEventNotFound       = EventError("Event not found")
	ErrEventsNotFound      = EventError("Events not found")
	ErrNothingFound        = EventError("Nothing found")
	ErrEventAlreadyExists  = EventError("Event already exists")
	ErrEventAlreadyDeleted = EventError("Event already deleted")
	//ErrOverlapping       = EventError("another event exists for this date")
	//ErrIncorrectDuration = EventError("duration is incorrect")
	//ErrDateBusy  = EventError("The date already busy")
)
