// Package nolimit implements cache with simple nolimit algorithm.
package nolimit

import (
	"sync"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// NoLimit is the base nolimit type.
type NoLimit struct {
	sync.RWMutex
	cache map[string]struct{}
}

// NewCache returns new cache struct.
func NewCache() (*NoLimit, error) {
	return &NoLimit{
		cache: make(map[string]struct{}),
	}, nil
}

// Clear clears all cache.
func (nl *NoLimit) Clear() {
	nl.cache = make(map[string]struct{})
}

// Add adds item into cache and returns deleted item.
func (nl *NoLimit) Add(item entities.CacheItem) (entities.CacheItem, error) {
	nl.RWMutex.Lock()
	defer nl.RWMutex.Unlock()
	nl.cache[item.Hash] = struct{}{}

	return entities.CacheItem{}, nil
}

// Get got item from cache.
func (nl *NoLimit) Get(hash string) (entities.CacheItem, bool) {
	if _, ok := nl.cache[hash]; !ok {
		return entities.CacheItem{}, false
	}

	return entities.CacheItem{}, true
}
