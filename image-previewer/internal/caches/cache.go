/*
 * Project: Image Previewer
 * Created on 26.01.2020 17:37
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package caches implements cacher interface.
package caches

import (
	"errors"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"

	"gitlab.com/tirava/image-previewer/internal/caches/nolimit"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/cache"
)

// NewCacher returns cache implementer.
func NewCacher(implementer string, storage storage.Storager) (cache.Cacher, error) {
	switch implementer {
	case "lru":
		//return lru.NewCache()
	case "lru2q":
		//return lru2q.NewCache()
	case "nolimit":
		return nolimit.NewCache(storage)
	}

	return nil, errors.New("incorrect cache implementer name")
}
