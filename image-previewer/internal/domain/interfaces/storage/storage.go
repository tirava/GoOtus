// Package storage implements storage interface.
package storage

import "gitlab.com/tirava/image-previewer/internal/domain/entities"

// Storager is the main interface for storage logic.
type Storager interface {
	Save(item entities.CacheItem) (bool, error)
	Load(hash string) (entities.CacheItem, error)
	Delete(item entities.CacheItem) error
	Close() error
	IsItemExist(hash string) (bool, string)
}
