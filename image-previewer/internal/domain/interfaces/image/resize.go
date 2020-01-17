/*
 * Project: Image Previewer
 * Created on 17.01.2020 11:01
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package image describes domain image interface.
package image

import (
	"image"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// Resizer is the main interface for resizing logic.
type Resizer interface {
	Resize(
		width, height uint,
		img image.Image,
		opts entities.ResizeOptions,
	) (image.Image, error)
}
