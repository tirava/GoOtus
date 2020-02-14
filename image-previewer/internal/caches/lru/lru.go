// Package lru implements LRU cache algorithm .
package lru

import (
	"sync"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/linkedlist"
)

// LRU is the base lru type.
type LRU struct {
	sync.RWMutex
	cache    map[string]*linkedlist.Element
	list     *linkedlist.List
	maxItems int
}

// NewCache returns new cache struct.
func NewCache(maxItems int) (*LRU, error) {
	return &LRU{
		cache:    make(map[string]*linkedlist.Element),
		list:     linkedlist.New(),
		maxItems: maxItems,
	}, nil
}

// Clear clears all cache.
func (l *LRU) Clear() {
	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()

	l.list = linkedlist.New()
	l.cache = make(map[string]*linkedlist.Element)
}

// Add adds item into cache and returns deleted item.
func (l *LRU) Add(item entities.CacheItem) (entities.CacheItem, error) {
	var deletedItem entities.CacheItem

	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()

	if len(l.cache) >= l.maxItems {
		deletedItem = *l.list.Back().Value().(*entities.CacheItem)
		l.list.Remove(l.list.Back())
		delete(l.cache, deletedItem.Hash)
	}

	l.list.PushFront(&item)
	l.cache[item.Hash] = l.list.Front()

	return deletedItem, nil
}

// Get got item from cache.
func (l *LRU) Get(hash string) (entities.CacheItem, bool) {
	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()

	elem, ok := l.cache[hash]
	if !ok {
		return entities.CacheItem{}, false
	}

	item := elem.Value().(*entities.CacheItem)
	l.list.Remove(elem)
	l.list.PushFront(item)
	l.cache[item.Hash] = l.list.Front()

	return *item, true
}
