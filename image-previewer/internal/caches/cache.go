// Package caches implements cacher interface.
package caches

import (
	"fmt"

	"gitlab.com/tirava/image-previewer/internal/caches/lru"
	"gitlab.com/tirava/image-previewer/internal/caches/nolimit"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/cache"
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"
	"gitlab.com/tirava/image-previewer/internal/models"
)

// NewCacher returns cache implementer.
func NewCacher(implementer string, storage storage.Storager, maxItems int) (cache.Cacher, error) {
	switch implementer {
	case models.LRU:
		return lru.NewCache(storage, maxItems)
	case models.NoLimit:
		return nolimit.NewCache(storage)
	}

	return nil, fmt.Errorf("incorrect cache implementer name: %s", implementer)
}
