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

	return l.storage.Save(item)
}

// Get got item from cache.
func (l *LRU) Get(hash string) (entities.CacheItem, bool, error) {
	if _, ok := l.cache[hash]; !ok {
		if ok, _ := l.storage.IsItemExist(hash); ok {
			item := entities.CacheItem{Hash: hash}
			if _, err := l.Add(item); err != nil {
				return entities.CacheItem{}, false, err
			}
		} else {
			return entities.CacheItem{}, false, nil
		}
	}

	item, err := l.storage.Load(hash)
	if err != nil {
		l.list.Remove(l.cache[hash])
		delete(l.cache, hash)

		return entities.CacheItem{}, false, err
	}

	l.moveItemToHead(l.cache[hash])

	return item, true, nil
}

func (l *LRU) moveItemToHead(item *list.Element) {
	data := item.Value.(*entities.CacheItem)

	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()
	l.list.Remove(item)

	l.list.PushFront(data)
	l.cache[data.Hash] = l.list.Front()
}

func (l *LRU) deleteLastItem() error {
	data := l.list.Back().Value.(*entities.CacheItem)

	l.RWMutex.Lock()
	l.list.Remove(l.list.Back())
	delete(l.cache, data.Hash)
	l.RWMutex.Unlock()

	if err := l.storage.Delete(*data); err != nil {
		return err
	}

	return nil
}
