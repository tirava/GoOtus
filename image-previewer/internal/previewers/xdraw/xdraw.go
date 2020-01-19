/*
 * Project: Image Previewer
 * Created on 19.01.2020 13:45
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package xdraw implements x/draw image resizer.
// golang.org/x/image/draw
package xdraw

import (
	"image"

	"golang.org/x/image/draw"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

const scaleFactor = 0.7

// XDraw is the base xdraw type.
type XDraw struct {
}

// NewXDraw returns new nfnt struct.
func NewXDraw() (*XDraw, error) {
	return &XDraw{}, nil
}

// Resize implements resize interface method.
func (n XDraw) Resize(
	width, height int, img image.Image,
	opts entities.ResizeOptions) image.Image {
	scaleX, scaleY := calcFactors(width, height, float64(img.Bounds().Dx()), float64(img.Bounds().Dy()))

	if width == 0 {
		width = int(scaleFactor + float64(img.Bounds().Dx())/scaleX)
	}

	if height == 0 {
		height = int(scaleFactor + float64(img.Bounds().Dy())/scaleY)
	}

	var ds draw.Interpolator

	switch opts.Interpolation {
	case entities.ApproxBiLinear:
		ds = draw.ApproxBiLinear
	case entities.Bilinear:
		ds = draw.BiLinear
	case entities.CatmullRom:
		ds = draw.CatmullRom
	default:
		ds = draw.NearestNeighbor
	}

	dr := image.Rect(0, 0, width, height)
	result := image.NewRGBA(dr)

	ds.Scale(result, dr, img, img.Bounds(), draw.Over, nil)

	return result
}

// Crop implements crop interface method.
func (n XDraw) Crop(
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

func calcFactors(width, height int, oldWidth, oldHeight float64) (scaleX, scaleY float64) {
	if width == 0 {
		if height == 0 {
			scaleX = 1.0
			scaleY = 1.0
		} else {
			scaleY = oldHeight / float64(height)
			scaleX = scaleY
		}
	} else {
		scaleX = oldWidth / float64(width)
		if height == 0 {
			scaleY = scaleX
		} else {
			scaleY = oldHeight / float64(height)
		}
	}

	return
}
