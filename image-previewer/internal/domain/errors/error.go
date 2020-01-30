/*
 * Project: Image Previewer
 * Created on 28.01.2020 15:46
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package errors implements base preview errors.
package errors

// PreviewError is the base type for preview errors.
type PreviewError string

// Error returns string error of the PreviewError.
func (e PreviewError) Error() string {
	return string(e)
}

// Errors.
var (
	ErrEmptyList             = PreviewError("cache list is empty")
	ErrItemNotFoundInStorage = PreviewError("item not found in storage")
)
