// Package preview implements business logic.
package preview

import (
	"fmt"
	"image"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/cache"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/encode"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/preview"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"
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
func NewPreview(prevImpl preview.Previewer, encImpl encode.Hasher, cacheImpl cache.Cacher, storImpl storage.Storager,
) (*Preview, error) {
	return &Preview{
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

	var wr, hr, x, y int

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

// CalcHash calculates hash for url.
func (p Preview) CalcHash(url string) (string, error) {
	hash, err := p.ImageURLEncoder.Encode(url)
	if err != nil {
		return "", err
	}

	return hash, nil
}

// IsItemInCache returns image if it in the cache.
func (p Preview) IsItemInCache(hash string) (entities.CacheItem, bool, error) {
	item, ok := p.Cacher.Get(hash)

	if !ok {
		if ok, _ := p.Storager.IsItemExist(hash); ok {
			item, err := p.Storager.Load(hash)
			if err != nil {
				return entities.CacheItem{}, false, err
			}

			if _, err := p.Cacher.Add(item); err != nil {
				return entities.CacheItem{}, false, err
			}

			return item, true, nil
		}

		return entities.CacheItem{}, false, nil
	}

	if item.Image == nil {
		var err error
		item, err = p.Storager.Load(hash)

		if err != nil {
			return entities.CacheItem{}, false, err
		}
	}

	return item, true, nil
}

// AddItemIntoCache adds image into cache.
func (p Preview) AddItemIntoCache(item entities.CacheItem) (bool, error) {
	delItem, err := p.Cacher.Add(item)
	if err != nil {
		return false, err
	}

	var errDel error
	if delItem.Hash != "" {
		errDel = p.Storager.Delete(delItem)
	}

	ok, errSave := p.Storager.Save(item)

	if errDel != nil || errSave != nil {
		err = fmt.Errorf("%s%s", errDel, errSave)
	}

	// no need raw bytes anymore
	item.RawBytes = nil

	return ok, err
}

// Close closes any open handlers.
func (p Preview) Close() error {
	return p.Storager.Close()
}
