package entities

import (
	"image"
)

// CacheItem is the base cache item.
type CacheItem struct {
	Image    image.Image
	ImgType  string
	Hash     string
	RawBytes []byte
}
