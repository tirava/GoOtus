/*
 * Project: Image Previewer
 * Created on 26.01.2020 20:27
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package inmemory implements storage in memory map.
package inmemory

import (
	"fmt"
	"sync"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
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
	im.storage[item.Hash] = item

	return false, nil
}

// Load loads item from the storage.
func (im *InMemory) Load(hash string) (entities.CacheItem, error) {
	item, ok := im.storage[hash]
	if !ok {
		return entities.CacheItem{}, fmt.Errorf("item not found in storage: %s", hash)
	}

	return item, nil
}

// Delete deletes item in the storage.
func (im *InMemory) Delete(hash string) error {
	im.RWMutex.Lock()
	defer im.RWMutex.Unlock()
	delete(im.storage, hash)

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
