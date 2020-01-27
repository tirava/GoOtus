/*
 * Project: Image Previewer
 * Created on 26.01.2020 20:19
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package storages implements storage interface.
package storages

import (
	"errors"

	"gitlab.com/tirava/image-previewer/internal/storages/disk"

	"gitlab.com/tirava/image-previewer/internal/storages/inmemory"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"
)

// NewStorager returns storage implementer.
func NewStorager(implementer, storPath string) (storage.Storager, error) {
	switch implementer {
	case "disk":
		return disk.NewStorage(storPath)
	case "inmemory":
		return inmemory.NewStorage()
	}

	return nil, errors.New("incorrect storage implementer name")
}
