// Package preview describes domain image interface.
package preview

import (
	"image"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// Previewer is the main interface for resizing logic.
type Previewer interface {
	Resize(
		width, height int, img image.Image,
		opts entities.ResizeOptions) image.Image
	Crop(
		width, height int, anchor image.Point,
		img image.Image) image.Image
}
