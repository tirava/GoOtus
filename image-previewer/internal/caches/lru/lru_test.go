package lru

import (
	"image"
	"testing"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

type cacheActions func(*testing.T, *LRU)

type expResult struct {
	ok    bool
	err   error
	order uint // from 1
}

const (
	maxItems = 2
	md5Fake  = "47dc34b1348a6b12d4b0fa5c350d08c4"
	md5Fake1 = "11111111111111111111111111111111"
	md5Fake2 = "22222222222222222222222222222222"
	md5Fake3 = "33333333333333333333333333333333"
	md5Fake4 = "44444444444444444444444444444444"
	md5Fake5 = "55555555555555555555555555555555"
	md5Fake6 = "66666666666666666666666666666666"
	md5Fake7 = "77777777777777777777777777777777"
)

// nolint:gochecknoglobals
var testImage = image.NewRGBA(image.Rect(0, 0, 100, 100))

// nolint:gochecknoglobals, gomnd
var testCasesLRU = []struct {
	description string
	actions     []cacheActions
}{
	{
		"Add item into cache, get it again",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5Fake,
			}, expResult{}),
			isItemInCache(md5Fake, expResult{ok: true}),
		},
	},
	{
		"Add item and clear all cache",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5Fake1,
			}, expResult{}),
			clearCache(),
			isItemInCache(md5Fake1, expResult{}),
		},
	},
	{
		"Add item into cache, old item deleted",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5Fake2,
			}, expResult{}),
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5Fake3,
			}, expResult{}),
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5Fake4,
			}, expResult{}),
			isItemInCache(md5Fake2, expResult{ok: false}),
		},
	},
	{
		"Get Last Recent Used item",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5Fake5,
			}, expResult{}),
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5Fake6,
			}, expResult{}),
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5Fake7,
			}, expResult{}),
			//checkItemOrder(md5Fake6, expResult{order: 2}),
			isItemInCache(md5Fake6, expResult{ok: true}),
			checkItemOrder(md5Fake6, expResult{order: 1}),
		},
	},
}

func addItemIntoCache(item entities.CacheItem, expected expResult) cacheActions {
	return func(t *testing.T, l *LRU) {
		_, err := l.Add(item)
		if err != nil && err.Error() != expected.err.Error() {
			t.Errorf("AddI() returned wrong, expected err=%s, got err=%s",
				expected.err, err)
		}
	}
}

func isItemInCache(hash string, expected expResult) cacheActions {
	return func(t *testing.T, l *LRU) {
		item, ok := l.Get(hash)
		if ok != expected.ok {
			t.Errorf("Get() returned wrong, expected ok=%t, got ok=%t",
				expected.ok, ok)
		}

		if ok && item.Hash != hash {
			t.Errorf("Get() returned wrong, expected hash=%s, got hash=%s",
				hash, item.Hash)
		}
	}
}

func clearCache() cacheActions {
	return func(t *testing.T, l *LRU) {
		l.Clear()
	}
}

func checkItemOrder(hash string, expected expResult) cacheActions {
	return func(t *testing.T, l *LRU) {
		var order uint
		for elem := l.list.Front(); elem != nil; elem = elem.Next() {
			order++

			if elem.Value().(*entities.CacheItem).Hash == hash {
				break
			}
		}

		if order != expected.order {
			t.Errorf("Item order wrong, expected order=%d, got order=%d",
				expected.order, order)
		}
	}
}

func TestCacheLRU(t *testing.T) {
	cache, _ := NewCache(maxItems)

	for _, test := range testCasesLRU {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			for _, act := range test.actions {
				act(t, cache)
			}
		})
	}
}
