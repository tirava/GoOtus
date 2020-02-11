// Package nfntcrop implements nfnt image resizer.
// https://github.com/nfnt/resize
// crop is implemented by Eugene Klimov
package nfntcrop

import (
	"image"
	"image/draw"

	"github.com/nfnt/resize"
	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// NfNtCrop is the base nfnt type.
type NfNtCrop struct {
}

// NewNfNtCrop returns new nfnt struct.
func NewNfNtCrop() (*NfNtCrop, error) {
	return &NfNtCrop{}, nil
}

// Resize implements resize interface method.
func (n NfNtCrop) Resize(
	width, height int, img image.Image,
	opts entities.ResizeOptions) image.Image {
	return resize.Resize(uint(width), uint(height), img,
		resize.InterpolationFunction(opts.Interpolation))
}

// Crop implements crop interface method.
func (n NfNtCrop) Crop(
	width, height int, anchor image.Point,
	img image.Image) image.Image {
	size := image.Point{X: width, Y: height}
	rMin := image.Point{X: anchor.X, Y: anchor.Y}
	cr := image.Rect(rMin.X, rMin.Y, rMin.X+size.X, rMin.Y+size.Y)
	cr = img.Bounds().Intersect(cr)

	result := image.NewRGBA(cr)
	draw.Draw(result, cr, img, cr.Min, draw.Src)

	return result
}
