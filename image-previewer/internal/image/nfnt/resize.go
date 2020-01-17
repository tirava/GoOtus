/*
 * Project: Image Previewer
 * Created on 17.01.2020 11:38
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package nfnt implements nfnt image resizer.
// https://github.com/nfnt/resize
package nfnt

import (
	"image"

	"github.com/nfnt/resize"
	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// NfNt is the base nfnt type.
type NfNt struct {
}

// NewNfNt returns new nfnt struct.
func NewNfNt() (*NfNt, error) {
	return &NfNt{}, nil
}

// Resize implements resize interface method.
func (n NfNt) Resize(
	width, height uint,
	img image.Image,
	opts entities.ResizeOptions,
) (image.Image, error) {
	return resize.Resize(width, height, img,
		resize.InterpolationFunction(opts.Interpolation)), nil
}
