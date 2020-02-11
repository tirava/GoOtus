// Package nolimit implements cache with simple nolimit algorithm.
package nolimit

import (
	"sync"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"
)

// NoLimit is the base nolimit type.
type NoLimit struct {
	sync.RWMutex
	cache   map[string]struct{}
	storage storage.Storager
}

// NewCache returns new cache struct.
func NewCache(storage storage.Storager) (*NoLimit, error) {
	return &NoLimit{
		cache:   make(map[string]struct{}),
		storage: storage,
	}, nil
}

// Clear clears all cache.
func (nl *NoLimit) Clear() {
	nl.cache = make(map[string]struct{})
}

// Add adds item into cache.
func (nl *NoLimit) Add(item entities.CacheItem) (bool, error) {
	nl.RWMutex.Lock()
	defer nl.RWMutex.Unlock()
	nl.cache[item.Hash] = struct{}{}

	return nl.storage.Save(item)
}

// Get got item from cache.
func (nl *NoLimit) Get(hash string) (entities.CacheItem, bool, error) {
	nl.RWMutex.Lock()
	_, ok := nl.cache[hash]
	nl.RWMutex.Unlock()

	if !ok {
		if ok, _ := nl.storage.IsItemExist(hash); ok {
			item := entities.CacheItem{Hash: hash}
			if _, err := nl.Add(item); err != nil {
				return entities.CacheItem{}, false, err
			}
		} else {
			return entities.CacheItem{}, false, nil
		}
	}

	item, err := nl.storage.Load(hash)
	if err != nil {
		return entities.CacheItem{}, false, err
	}

	return item, true, nil
}
