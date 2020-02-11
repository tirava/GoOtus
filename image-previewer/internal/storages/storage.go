// Package storages implements storage interface.
package storages

import (
	"fmt"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/storage"
	"gitlab.com/tirava/image-previewer/internal/models"
	"gitlab.com/tirava/image-previewer/internal/storages/disk"
	"gitlab.com/tirava/image-previewer/internal/storages/inmemory"
)

// NewStorager returns storage implementer.
func NewStorager(implementer, storPath string) (storage.Storager, error) {
	switch implementer {
	case models.Disk:
		return disk.NewStorage(storPath)
	case models.InMemory:
		return inmemory.NewStorage()
	}

	return nil, fmt.Errorf("incorrect storage implementer name: %s", implementer)
}
