/*
 * Project: Image Previewer
 * Created on 17.01.2020 11:01
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package entities implements various options models.
package entities

type interpolation int

// Constants for interpolation method.
const (
	NearestNeighbor interpolation = iota
	Bilinear
	Bicubic
	MitchellNetravali
	Lanczos2
	Lanczos3
)

// ResizeOptions is the base resize options.
type ResizeOptions struct {
	Interpolation interpolation
}
