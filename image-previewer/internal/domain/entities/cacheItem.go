/*
 * Project: Image Previewer
 * Created on 26.01.2020 14:12
 * Copyright (c) 2020 - Eugene Klimov
 */

package entities

import (
	"image"
)

// CacheItem is the base cache item.
type CacheItem struct {
	Image    image.Image
	ImgType  string
	Hash     string
	StorPath string
	RawBytes []byte
}
