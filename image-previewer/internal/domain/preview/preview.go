/*
 * Project: Image Previewer
 * Created on 17.01.2020 11:57
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package preview implements business logic.
package preview

import (
	"image"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/preview"
	"gitlab.com/tirava/image-previewer/internal/previewers"
)

const anchorBaseDivImage = 2 // 2 - center

// Preview is the main Preview struct.
type Preview struct {
	Previewer preview.Previewer
}

// NewPreview inits main Preview fields.
func NewPreview(name string) (Preview, error) {
	p, err := previewers.NewPreviewer(name)
	if err != nil {
		return Preview{}, err
	}

	return Preview{Previewer: p}, nil
}

// Preview returns preview result image.
func (p Preview) Preview(width, height int, img image.Image, opts entities.ResizeOptions) image.Image {
	if width < 0 || height < 0 || (width == 0 && height == 0) {
		return img
	}

	wr, hr := 0, 0
	x, y := 0, 0

	wb := float64(width) / float64(img.Bounds().Max.X)
	hb := float64(height) / float64(img.Bounds().Max.Y)

	if wb > hb {
		wr = width
	} else {
		hr = height
	}

	pr := p.Previewer.Resize(wr, hr, img, opts)

	if width > height {
		y = (pr.Bounds().Max.Y - height) / anchorBaseDivImage
	} else {
		x = (pr.Bounds().Max.X - width) / anchorBaseDivImage
	}

	pr = p.Previewer.Crop(width, height, image.Point{X: x, Y: y}, pr)

	return pr
}
