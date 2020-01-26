/*
 * Project: Image Previewer
 * Created on 26.01.2020 17:44
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package nolimit implements cache with simple nolimit algorithm.
package nolimit

import (
	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"
)

// NoLimit is the base nolimit type.
type NoLimit struct {
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

// Add adds item into cache.
func (nl NoLimit) Add(item entities.CacheItem) error {
	nl.cache[item.Hash] = struct{}{}

	return nl.storage.Save(item)
}

// Get got item from cache.
func (nl NoLimit) Get(hash string) (entities.CacheItem, bool, error) {
	if _, ok := nl.cache[hash]; !ok {
		return entities.CacheItem{}, false, nil
	}

	item, err := nl.storage.Load(hash)
	if err != nil {
		return entities.CacheItem{}, false, err
	}

	return item, true, nil
}
