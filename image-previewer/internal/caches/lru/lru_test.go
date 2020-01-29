/*
 * Project: Image Previewer
 * Created on 28.01.2020 22:33
 * Copyright (c) 2020 - Eugene Klimov
 */

package lru

import (
	"fmt"
	"image"
	"testing"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"

	"gitlab.com/tirava/image-previewer/internal/domain/preview"
	"gitlab.com/tirava/image-previewer/internal/encoders"
	"gitlab.com/tirava/image-previewer/internal/previewers"
	"gitlab.com/tirava/image-previewer/internal/storages"
)

const (
	fakeURL     = "http://fake/image.tiff"
	md5FakeURL  = "47dc34b1348a6b12d4b0fa5c350d08c4"
	md5FakeURL1 = "11111111111111111111111111111111"
	md5FakeURL2 = "22222222222222222222222222222222"
)

type cacheActions func(*testing.T, preview.Preview)

type expResult struct {
	//item entities.CacheItem
	ok  bool
	err error
}

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
				Image:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
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
				Image:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				ImgType: "gif",
				Hash:    md5FakeURL,
			}, expResult{}),
			addItemIntoCache(entities.CacheItem{
				Image:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				ImgType: "gif",
				Hash:    md5FakeURL1,
			}, expResult{}),
			addItemIntoCache(entities.CacheItem{
				Image:   image.NewRGBA(image.Rect(0, 0, 100, 100)),
				ImgType: "gif",
				Hash:    md5FakeURL2,
			}, expResult{}),
			isItemInCache(fakeURL, expResult{ok: false}),
		},
		2,
	},
	{
		"Get absent item from cache",
		[]cacheActions{
			isItemInCache("http://fake/image.tiff", expResult{}),
		},
		2,
	},
	//{
	//	//"Get presented item from cache",
	//},
	//{
	//	//"Get absent item from cache presented in storage",
	//},
}

func addItemIntoCache(item entities.CacheItem, expected expResult) cacheActions {
	return func(t *testing.T, p preview.Preview) {
		_, err := p.AddItemIntoCache(item)
		if err != expected.err {
			t.Errorf("AddItemIntoCache() returned wrong, expected err=%s, got err=%s",
				expected.err, err)
		}
	}
}

func isItemInCache(url string, expected expResult) cacheActions {
	return func(t *testing.T, p preview.Preview) {
		item, ok, err := p.IsItemInCache(url)
		if ok != expected.ok {
			t.Errorf("IsItemInCache() returned wrong, expected ok=%t, got ok=%t",
				expected.ok, ok)
		}

		if err != expected.err {
			t.Errorf("IsItemInCache() returned wrong, expected err=%s, got err=%s",
				expected.err, err)
		}

		fmt.Println(item)
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
	maxItems int) (preview.Preview, error) {
	prev, err := previewers.NewPreviewer(prevImpl)
	if err != nil {
		return preview.Preview{}, err
	}

	enc, err := encoders.NewImageURLEncoder(encImpl)
	if err != nil {
		return preview.Preview{}, err
	}

	stor, err := storages.NewStorager(storImpl, storPath)
	if err != nil {
		return preview.Preview{}, err
	}

	cash, err := NewCache(stor, maxItems)
	if err != nil {
		return preview.Preview{}, err
	}

	return preview.NewPreview(prev, enc, cash, stor)
}
