/*
 * Project: Image Previewer
 * Created on 28.01.2020 14:41
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package lru implements LRU cache algorithm .
package lru

import (
	"container/list"
	"sync"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"
)

// LRU is the base lru type.
type LRU struct {
	sync.RWMutex
	cache    map[string]*list.Element
	list     *list.List
	maxItems int
	storage  storage.Storager
}

// NewCache returns new cache struct.
func NewCache(storage storage.Storager, maxItems int) (*LRU, error) {
	return &LRU{
		cache:    make(map[string]*list.Element),
		list:     list.New(),
		maxItems: maxItems,
		storage:  storage,
	}, nil
}

// Clear clears all cache.
func (l *LRU) Clear() {
	l.list = list.New()
	l.cache = make(map[string]*list.Element)
}

// Add adds item into cache.
func (l *LRU) Add(item entities.CacheItem) (bool, error) {
	if len(l.cache) >= l.maxItems {
		if err := l.deleteLastItem(); err != nil {
			return false, err
		}
	}

	l.RWMutex.Lock()
	l.list.PushFront(&item)
	l.cache[item.Hash] = l.list.Front()
	l.RWMutex.Unlock()

	ok, err := l.storage.Save(item)

	// no need raw bytes and image in the lru list
	item.RawBytes = nil

	return ok, err
}

// Get got item from cache.
func (l *LRU) Get(hash string) (entities.CacheItem, bool, error) {
	l.RWMutex.Lock()
	elem, ok := l.cache[hash]
	l.RWMutex.Unlock()

	if !ok {
		if ok, _ := l.storage.IsItemExist(hash); ok {
			item, err := l.storage.Load(hash)
			if err != nil {
				return entities.CacheItem{}, false, err
			}

			if _, err := l.Add(item); err != nil {
				return entities.CacheItem{}, false, err
			}

			return item, true, nil
		}

		return entities.CacheItem{}, false, nil
	}

	item := elem.Value.(*entities.CacheItem)
	l.moveItemToHead(elem)

	return *item, true, nil
}

func (l *LRU) moveItemToHead(item *list.Element) {
	if item == nil {
		return
	}

	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()

	data := item.Value.(*entities.CacheItem)

	l.list.Remove(item)

	l.list.PushFront(data)
	l.cache[data.Hash] = l.list.Front()
}

func (l *LRU) deleteLastItem() error {
	l.RWMutex.Lock()
	data := l.list.Back().Value.(*entities.CacheItem)

	l.list.Remove(l.list.Back())
	delete(l.cache, data.Hash)
	l.RWMutex.Unlock()

	if err := l.storage.Delete(*data); err != nil {
		return err
	}

	return nil
}
