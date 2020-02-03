/*
 * Project: Image Previewer
 * Created on 28.01.2020 22:33
 * Copyright (c) 2020 - Eugene Klimov
 */

package lru

import (
	"image"
	"testing"

	"gitlab.com/tirava/image-previewer/internal/domain/errors"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"

	"gitlab.com/tirava/image-previewer/internal/domain/preview"
	"gitlab.com/tirava/image-previewer/internal/encoders"
	"gitlab.com/tirava/image-previewer/internal/previewers"
	"gitlab.com/tirava/image-previewer/internal/storages"
)

const (
	fakeURL          = "http://fake/image.tiff"
	md5FakeURL       = "47dc34b1348a6b12d4b0fa5c350d08c4"
	md5FakeURL1      = "11111111111111111111111111111111"
	md5FakeURL2      = "22222222222222222222222222222222"
	md5DeleteFakeURL = "9201dafe08a33bbb90680a051adde096"
	md5LoadFakeURL   = "49f351f3016db4e5f00dd2eb683f56b3"
)

type cacheActions func(*testing.T, *preview.Preview)

type expResult struct {
	ok  bool
	err error
}

// nolint
var testImage = image.NewRGBA(image.Rect(0, 0, 100, 100))

// nolint
var testCases = []struct {
	description string
	actions     []cacheActions
	maxItems    int
}{
	{
		"Add item into cache, get it again",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL,
			}, expResult{}),
			isItemInCache(fakeURL, expResult{ok: true}),
		},
		1,
	},
	{
		"Add item into cache, old item deleted",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL,
			}, expResult{}),
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL1,
			}, expResult{}),
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL2,
			}, expResult{}),
			isItemInCache(fakeURL, expResult{ok: false}),
		},
		2,
	},
	{
		"Item absent in cache, it in storage but error loading",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5LoadFakeURL,
			}, expResult{}),
			clearCache(),
			isItemInCache("*testing.T.Load", expResult{err: errors.ErrItemNotFoundInStorage}),
		},
		1,
	},
	{
		"Get absent item from cache",
		[]cacheActions{
			isItemInCache("http://fake/image.tiff", expResult{}),
		},
		2,
	},
	{
		"Get absent item in cache but it presented in storage",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL,
			}, expResult{}),
			clearCache(),
			isItemInCache("http://fake/image.tiff", expResult{ok: true}),
		},
		1,
	},
	{
		"Get absent item in cache, it in storage but error adding into cache",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5DeleteFakeURL,
			}, expResult{}),
			clearCache(),
			isItemInCache("*testing.T", expResult{ok: true}),
			addItemIntoStorage(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL,
			}),
			isItemInCache("http://fake/image.tiff", expResult{err: errors.ErrItemNotFoundInStorage}),
		},
		1,
	},
}

func clearCache() cacheActions {
	return func(t *testing.T, p *preview.Preview) {
		p.Cacher.Clear()
	}
}

func addItemIntoCache(item entities.CacheItem, expected expResult) cacheActions {
	return func(t *testing.T, p *preview.Preview) {
		_, err := p.AddItemIntoCache(item)
		if err != expected.err {
			t.Errorf("AddItemIntoCache() returned wrong, expected err=%s, got err=%s",
				expected.err, err)
		}
	}
}

func addItemIntoStorage(item entities.CacheItem) cacheActions {
	return func(t *testing.T, p *preview.Preview) {
		_, _ = p.Storager.Save(item)
	}
}

// nolint
func isItemInCache(url string, expected expResult) cacheActions {
	return func(t *testing.T, p *preview.Preview) {
		_, ok, err := p.IsItemInCache(url)
		if ok != expected.ok {
			t.Errorf("IsItemInCache() returned wrong, expected ok=%t, got ok=%t",
				expected.ok, ok)
		}

		if err != nil && err.Error() == expected.err.Error() {
			err = expected.err
		}

		if err != expected.err {
			t.Errorf("IsItemInCache() returned wrong, expected err=%s, got err=%s",
				expected.err, err)
		}
	}
}

// nolint
func deleteItemInCache(item entities.CacheItem, expected expResult) cacheActions {
	return func(t *testing.T, p *preview.Preview) {
		err := p.Storager.Delete(item)
		if err != expected.err {
			t.Errorf("Storager.Delete() returned wrong, expected err=%s, got err=%s",
				expected.err, err)
		}
	}
}

// nolint
func TestCache(t *testing.T) {
	for _, test := range testCases {
		prev, err := initPreview(
			"xdraw", "md5", "inmemory", "", test.maxItems)
		if err != nil {
			t.Fatal(err)
		}

		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			for _, act := range test.actions {
				act(t, prev)
			}
		})
	}
}

func initPreview(prevImpl, encImpl, storImpl, storPath string,
	maxItems int) (*preview.Preview, error) {
	prev, err := previewers.NewPreviewer(prevImpl)
	if err != nil {
		return nil, err
	}

	enc, err := encoders.NewImageURLEncoder(encImpl)
	if err != nil {
		return nil, err
	}

	stor, err := storages.NewStorager(storImpl, storPath)
	if err != nil {
		return nil, err
	}

	cash, err := NewCache(stor, maxItems)
	if err != nil {
		return nil, err
	}

	return preview.NewPreview(prev, enc, cash, stor)
}
