/*
 * Project: Image Previewer
 * Created on 26.01.2020 20:27
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package inmemory implements storage in memory map.
package inmemory

import (
	"fmt"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// InMemory is the base inmemory type.
type InMemory struct {
	storage map[string]entities.CacheItem
}

// NewStorage returns new storage struct.
func NewStorage() (*InMemory, error) {
	return &InMemory{
		storage: make(map[string]entities.CacheItem),
	}, nil
}

// Save saves item in the storage.
func (im InMemory) Save(item entities.CacheItem) error {
	im.storage[item.Hash] = item

	return nil
}

// Load loads item from the storage.
func (im InMemory) Load(hash string) (entities.CacheItem, error) {
	item, ok := im.storage[hash]
	if !ok {
		return entities.CacheItem{}, fmt.Errorf("item not found in storage: %s", hash)
	}

	return item, nil
}

// Delete deletes item in the storage.
func (im InMemory) Delete(hash string) error {
	return nil
}
