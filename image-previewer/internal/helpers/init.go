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
	"gitlab.com/tirava/image-previewer/internal/models"
	"gitlab.com/tirava/image-previewer/internal/previewers"
	"gitlab.com/tirava/image-previewer/internal/storages"
)

// InitPreview returns Preview with implementors
func InitPreview(conf models.Config) (*preview.Preview, error) {
	prev, err := previewers.NewPreviewer(conf.Previewer)
	if err != nil {
		return nil, err
	}

	enc, err := encoders.NewImageURLEncoder(conf.ImageURLEncoder)
	if err != nil {
		return nil, err
	}

	stor, err := storages.NewStorager(conf.Storager, conf.StoragePath)
	if err != nil {
		return nil, err
	}

	cash, err := caches.NewCacher(conf.Cacher, stor, conf.MaxCacheItems)
	if err != nil {
		return nil, err
	}

	return preview.NewPreview(prev, enc, cash, stor)
}
