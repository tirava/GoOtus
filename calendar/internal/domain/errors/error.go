/*
 * HomeWork-8: Calendar protobuf preparation
 * Created on 28.10.2019 22:36
 * Copyright (c) 2019 - Eugene Klimov
 */

// Package errors implements base calendar errors.
package errors

import "errors"

// Errors
var (
	ErrEventNotFound       = errors.New("event not found")
	ErrEventAlreadyDeleted = errors.New("event already deleted")
)
