// Package lru implements LRU cache algorithm .
package lru

import (
	"container/list"
	"sync"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// LRU is the base lru type.
type LRU struct {
	sync.RWMutex
	cache    map[string]*list.Element
	list     *list.List
	maxItems int
}

// NewCache returns new cache struct.
func NewCache(maxItems int) (*LRU, error) {
	return &LRU{
		cache:    make(map[string]*list.Element),
		list:     list.New(),
		maxItems: maxItems,
	}, nil
}

// Clear clears all cache.
func (l *LRU) Clear() {
	l.list = list.New()
	l.cache = make(map[string]*list.Element)
}

// Add adds item into cache and returns deleted item.
func (l *LRU) Add(item entities.CacheItem) (entities.CacheItem, error) {
	deletedItem := entities.CacheItem{}

	if len(l.cache) >= l.maxItems {
		deletedItem = l.deleteLastItem()
	}

	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()

	l.list.PushFront(&item)
	l.cache[item.Hash] = l.list.Front()

	return deletedItem, nil
}

// Get got item from cache.
func (l *LRU) Get(hash string) (entities.CacheItem, bool) {
	l.RWMutex.Lock()
	elem, ok := l.cache[hash]
	l.RWMutex.Unlock()

	if !ok {
		return entities.CacheItem{}, false
	}

	item := elem.Value.(*entities.CacheItem)
	l.moveItemToHead(elem)

	return *item, true
}

func (l *LRU) moveItemToHead(item *list.Element) {
	l.RWMutex.Lock()
	defer l.RWMutex.Unlock()

	data := item.Value.(*entities.CacheItem)

	l.list.Remove(item)

	l.list.PushFront(data)
	l.cache[data.Hash] = l.list.Front()
}

func (l *LRU) deleteLastItem() entities.CacheItem {
	l.RWMutex.Lock()
	data := l.list.Back().Value.(*entities.CacheItem)

	l.list.Remove(l.list.Back())
	delete(l.cache, data.Hash)
	l.RWMutex.Unlock()

	return *data
}
