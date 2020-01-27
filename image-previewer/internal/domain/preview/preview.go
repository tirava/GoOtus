/*
 * Project: Image Previewer
 * Created on 17.01.2020 11:57
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package preview implements business logic.
package preview

import (
	"image"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/cache"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/encode"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/preview"
)

const anchorBaseDivImage = 2 // 2 - center

// Preview is the main Preview struct.
type Preview struct {
	Previewer       preview.Previewer
	ImageURLEncoder encode.Hasher
	Cacher          cache.Cacher
	Storager        storage.Storager
}

// NewPreview inits main Preview fields.
func NewPreview(
	prevImpl preview.Previewer,
	encImpl encode.Hasher,
	cacheImpl cache.Cacher,
	storImpl storage.Storager,
) (Preview, error) {
	return Preview{
		Previewer:       prevImpl,
		ImageURLEncoder: encImpl,
		Cacher:          cacheImpl,
		Storager:        storImpl,
	}, nil
}

// Preview returns preview result image.
func (p Preview) Preview(width, height int, img image.Image, opts entities.ResizeOptions) image.Image {
	if width <= 0 || height <= 0 {
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

// IsItemInCache returns image if it in the cache.
func (p Preview) IsItemInCache(url string) (entities.CacheItem, bool, error) {
	hash, err := p.ImageURLEncoder.Encode(url)
	if err != nil {
		return entities.CacheItem{}, false, err
	}

	item, ok, err := p.Cacher.Get(hash)
	if !ok {
		return entities.CacheItem{
			Hash: hash,
		}, false, err
	}

	return item, true, nil
}

// AddItemIntoCache adds image into cache.
func (p Preview) AddItemIntoCache(item entities.CacheItem) (bool, error) {
	return p.Cacher.Add(item)
}

// Close closes any open handlers.
func (p Preview) Close() error {
	return p.Storager.Close()
}
