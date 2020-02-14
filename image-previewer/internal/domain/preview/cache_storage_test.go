package preview

import (
	"fmt"
	"image"
	"testing"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
	"gitlab.com/tirava/image-previewer/internal/domain/errors"
	"gitlab.com/tirava/image-previewer/internal/models"
)

const (
	maxItems         = 2
	fakeURL          = "http://fake/image.tiff"
	md5FakeURL       = "47dc34b1348a6b12d4b0fa5c350d08c4"
	md5FakeURL1      = "11111111111111111111111111111111"
	md5FakeURL2      = "22222222222222222222222222222222"
	deleteFakeURL    = "*testing.T"
	md5DeleteFakeURL = "9201dafe08a33bbb90680a051adde096"
	loadFakeURL      = "*testing.T.Load"
	md5LoadFakeURL   = "49f351f3016db4e5f00dd2eb683f56b3"
	loadFakeURL1     = "*testing.T.Load1"
	md5LoadFakeURL1  = "34db0fee103468f69d272ad042b43e86"
	md5SaveFakeURL   = "9d14351378dd8a31fb09ffac6c71ca6e"
	fakeURL3         = "fakeURL3"
	md5FakeURL3      = "a84f60db95707440f22a395771bdc42c"
	fakeURL4         = "fakeURL4"
	md5FakeURL4      = "eea58bcbfbe00a312a5682fc8f5cac59"
	fakeURL5         = "fakeURL5"
	md5FakeURL5      = "6a45921352bf387570c5b53e2e0e5e40"
	fakeURL6         = "fakeURL6"
	md5FakeURL6      = "3562fd37273fb46142362fde5d52a085"
)

type cacheActions func(*testing.T, *Preview)

type expResult struct {
	ok  bool
	err error
}

// nolint:gochecknoglobals
var testImage = image.NewRGBA(image.Rect(0, 0, 100, 100))

// nolint:gochecknoglobals
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
				Hash:    md5FakeURL,
			}, expResult{}),
			isItemInCache(fakeURL, expResult{ok: true}),
		},
	},
	{
		"Add item into cache, old item deleted",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL3,
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
			isItemInCache(fakeURL3, expResult{ok: false}),
		},
	},
	{
		"Get absent item from cache",
		[]cacheActions{
			isItemInCache("jvekvekjvbekjvbekvb", expResult{}),
		},
	},
	{
		"Get absent item in cache but it presented in storage",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL4,
			}, expResult{}),
			clearCache(),
			isItemInCache(fakeURL4, expResult{ok: true}),
		},
	},
	{
		"Get absent item in cache, it in storage but error adding into cache",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5DeleteFakeURL,
			}, expResult{}),
			isItemInCache(deleteFakeURL, expResult{ok: true}),
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL5,
			}, expResult{}),
			isItemInCache(fakeURL5, expResult{ok: true}),
			addItemIntoStorage(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5FakeURL6,
			}),
			isItemInCache(fakeURL6, expResult{err: errors.ErrItemNotFoundInStorage}),
		},
	},
}

// nolint:gochecknoglobals
var testCasesNoLimit = []struct {
	description string
	actions     []cacheActions
}{
	{
		"Add item into cache, get it again but error loading from storage",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5LoadFakeURL,
			}, expResult{}),
			isItemInCache(loadFakeURL, expResult{err: errors.ErrItemNotFoundInStorage}),
		},
	},
	{
		"Item absent in cache, it in storage but error loading",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5LoadFakeURL1,
			}, expResult{}),
			clearCache(),
			isItemInCache(loadFakeURL1, expResult{err: errors.ErrItemNotFoundInStorage}),
		},
	},
	{
		"Add item into cache but error save",
		[]cacheActions{
			addItemIntoCache(entities.CacheItem{
				Image:   testImage,
				ImgType: "gif",
				Hash:    md5SaveFakeURL,
			}, expResult{err: fmt.Errorf("%s %s", *new(error), errors.ErrSaveIntoStorage)}),
		},
	},
}

func clearCache() cacheActions {
	return func(t *testing.T, p *Preview) {
		p.Cacher.Clear()
	}
}

func addItemIntoCache(item entities.CacheItem, expected expResult) cacheActions {
	return func(t *testing.T, p *Preview) {
		defer func() {
			if r := recover(); r != nil {
				t.Log("Recovered in addItemIntoCache()", r)
			}
		}()

		_, err := p.AddItemIntoCache(item)
		if err != nil && err.Error() != expected.err.Error() {
			t.Errorf("AddItemIntoCache() returned wrong, expected err=%s, got err=%s",
				expected.err, err)
		}
	}
}

func addItemIntoStorage(item entities.CacheItem) cacheActions {
	return func(t *testing.T, p *Preview) {
		_, _ = p.Storager.Save(item)
	}
}

// nolint:unparam
func isItemInCache(url string, expected expResult) cacheActions {
	return func(t *testing.T, p *Preview) {
		hash, err := p.CalcHash(url)
		//fmt.Println(url, hash)
		if err != nil {
			t.Errorf("CalcHash() returned error: %s", err)
		}

		_, ok, err := p.IsItemInCache(hash)
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

// nolint:unused, deadcode
func deleteItemInCache(item entities.CacheItem, expected expResult) cacheActions {
	return func(t *testing.T, p *Preview) {
		err := p.Storager.Delete(item)
		if err != expected.err {
			t.Errorf("Storager.Delete() returned wrong, expected err=%s, got err=%s",
				expected.err, err)
		}
	}
}

func TestLRUCacheStorage(t *testing.T) {
	conf := models.Config{
		Previewer:       "xdraw",
		ImageURLEncoder: "md5",
		Cacher:          "lru",
		MaxCacheItems:   maxItems,
		Storager:        "inmemory",
	}

	prev, err := initPreview(conf)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range testCasesLRU {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			for _, act := range test.actions {
				act(t, prev)
			}
		})
	}
}

func TestNoLimitCacheStorage(t *testing.T) {
	conf := models.Config{
		Previewer:       "xdraw",
		ImageURLEncoder: "md5",
		Cacher:          "nolimit",
		MaxCacheItems:   maxItems,
		Storager:        "inmemory",
	}

	prev, err := initPreview(conf)
	if err != nil {
		t.Fatal(err)
	}

	for _, test := range testCasesNoLimit {
		test := test
		t.Run(test.description, func(t *testing.T) {
			t.Parallel()
			for _, act := range test.actions {
				act(t, prev)
			}
		})
	}
}
