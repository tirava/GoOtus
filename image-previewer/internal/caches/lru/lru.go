/*
 * Project: Image Previewer
 * Created on 28.01.2020 14:41
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package lru implements LRU cache algorithm .
package lru

import (
	"container/list"
	"fmt"
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

// Add adds item into cache.
func (l *LRU) Add(item entities.CacheItem) (bool, error) {
	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()

	if len(l.cache) >= l.maxItems {
		if err := l.deleteLastItem(); err != nil {
			return false, err
		}
	}

	l.list.PushFront(&item)
	l.cache[item.Hash] = l.list.Front()

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
		if errr := l.list.Remove(l.cache[hash]); errr != nil {
			return entities.CacheItem{}, false, fmt.Errorf("%w - %s", err, errr)
		}

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
	defer l.RWMutex.Unlock()
	l.list.Remove(l.list.Back())

	delete(l.cache, data.Hash)

	if err := l.storage.Delete(*data); err != nil {
		return err
	}

	return nil
}

// nolint
func (l *LRU) debugList() string {
	s, i := "", 0
	for e := l.list.Front(); e != nil; e = e.Next() {
		i++

		data := e.Value.(*entities.CacheItem)
		s += fmt.Sprintf("%d: %s\n", i, data.Hash)
	}

	return s
}
