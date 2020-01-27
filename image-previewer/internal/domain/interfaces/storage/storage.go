/*
 * Project: Image Previewer
 * Created on 26.01.2020 14:59
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package storage implements storage interface.
package storage

import "gitlab.com/tirava/image-previewer/internal/domain/entities"

// Storager is the main interface for storage logic.
type Storager interface {
	Save(item entities.CacheItem) (bool, error)
	Load(hash string) (entities.CacheItem, error)
	Delete(hash string) error
	Close() error
	IsItemExist(hash string) (bool, string)
}
