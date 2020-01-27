/*
 * Project: Image Previewer
 * Created on 27.01.2020 12:24
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package helpers implements helpers funcs.
package helpers

import (
	"gitlab.com/tirava/image-previewer/internal/caches"
	"gitlab.com/tirava/image-previewer/internal/domain/preview"
	"gitlab.com/tirava/image-previewer/internal/encoders"
	"gitlab.com/tirava/image-previewer/internal/previewers"
	"gitlab.com/tirava/image-previewer/internal/storages"
)

// InitPreview returns Preview with implementors
func InitPreview(prevImpl, encImpl, cacheImpl, storImpl, storPath string) (preview.Preview, error) {
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

	cash, err := caches.NewCacher(cacheImpl, stor)
	if err != nil {
		return preview.Preview{}, err
	}

	return preview.NewPreview(prev, enc, cash, stor)
}
