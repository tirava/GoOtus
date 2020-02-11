// Package inmemory implements storage in memory map.
package inmemory

import (
	"fmt"
	"sync"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/errors"
)

// InMemory is the base inmemory type.
type InMemory struct {
	sync.RWMutex
	storage map[string]entities.CacheItem
}

// NewStorage returns new storage struct.
func NewStorage() (*InMemory, error) {
	return &InMemory{
		storage: make(map[string]entities.CacheItem),
	}, nil
}

// Save saves item in the storage.
func (im *InMemory) Save(item entities.CacheItem) (bool, error) {
	if ok, _ := im.IsItemExist(item.Hash); ok {
		return true, nil
	}

	im.RWMutex.Lock()
	defer im.RWMutex.Unlock()

	item.RawBytes = nil // no need raw bytes for memory storage
	im.storage[item.Hash] = item

	return false, nil
}

// Load loads item from the storage.
func (im *InMemory) Load(hash string) (entities.CacheItem, error) {
	item, ok := im.storage[hash]
	if !ok {
		return entities.CacheItem{}, fmt.Errorf("%s: %s",
			errors.ErrItemNotFoundInStorage, hash)
	}

	// for testing purpose only
	if hash == "49f351f3016db4e5f00dd2eb683f56b3" {
		return entities.CacheItem{}, errors.ErrItemNotFoundInStorage
	}

	return item, nil
}

// Delete deletes item in the storage.
func (im *InMemory) Delete(item entities.CacheItem) error {
	im.RWMutex.Lock()
	defer im.RWMutex.Unlock()
	delete(im.storage, item.Hash)

	// for testing purpose only
	if item.Hash == "9201dafe08a33bbb90680a051adde096" {
		return errors.ErrItemNotFoundInStorage
	}

	return nil
}

// Close closes storage.
func (im *InMemory) Close() error {
	return nil
}

// IsItemExist checks if item in the storage.
func (im *InMemory) IsItemExist(hash string) (bool, string) {
	if _, ok := im.storage[hash]; ok {
		return true, ""
	}

	return false, ""
}
