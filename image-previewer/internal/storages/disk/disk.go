// Package disk implements storage on disk.
package disk

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/tirava/image-previewer/internal/domain/entities"
)

// Disk is the base disk type.
type Disk struct {
	storPath string
	imgTypes []string
}

// NewStorage returns new storage struct.
func NewStorage(storPath string) (*Disk, error) {
	ns := &Disk{
		storPath: storPath,
		imgTypes: []string{"jpeg", "png", "gif"},
	}

	if err := os.MkdirAll(storPath, 0711); err != nil {
		return ns, fmt.Errorf("unable to create cache dir '%s': %w",
			storPath, err)
	}

	return ns, nil
}

// Save saves item in the storage.
func (d *Disk) Save(item entities.CacheItem) (bool, error) {
	fileName := d.buildFileName(item.Hash, item.ImgType)

	if ok, _ := d.IsItemExist(item.Hash); ok {
		return true, nil
	}

	outFile, err := os.Create(fileName)
	if err != nil {
		return false, fmt.Errorf("unable to create cache file '%s' in '%s': %w",
			fileName, d.storPath, err)
	}
	defer outFile.Close()

	if _, err = outFile.Write(item.RawBytes); err != nil {
		return false, fmt.Errorf("unable to write cache file '%s': %w", fileName, err)
	}

	return false, nil
}

// Load loads item from the storage.
func (d *Disk) Load(hash string) (entities.CacheItem, error) {
	ok, fileName := d.IsItemExist(hash)
	if !ok {
		return entities.CacheItem{}, fmt.Errorf("cache item not found while loading: %s", hash)
	}

	file, err := os.Open(fileName)
	if err != nil {
		return entities.CacheItem{}, fmt.Errorf("can't open cache item file '%s': %w", fileName, err)
	}
	defer file.Close()

	var img image.Image

	ext := filepath.Ext(fileName)
	ext = strings.TrimLeft(ext, ".")

	switch ext {
	case "jpeg":
		img, err = jpeg.Decode(file)
	case "png":
		img, err = png.Decode(file)
	case "gif":
		img, err = gif.Decode(file)
	}

	if err != nil {
		return entities.CacheItem{}, fmt.Errorf("can't decode cache item file '%s': %w", fileName, err)
	}

	item := entities.CacheItem{
		Image:   img,
		ImgType: ext,
		Hash:    hash,
	}

	return item, nil
}

// Delete deletes item in the storage.
func (d *Disk) Delete(item entities.CacheItem) error {
	ok, fileName := d.IsItemExist(item.Hash)
	if !ok {
		return fmt.Errorf("cache item not found while delete: %s", item.Hash)
	}

	if err := os.Remove(fileName); err != nil {
		return fmt.Errorf("can't delete cache item file '%s': %w", fileName, err)
	}

	return nil
}

// Close closes storage and removes cached files.
func (d *Disk) Close() error {
	return nil // save cache for next run but may be return os.RemoveAll(d.storPath)
}

// IsItemExist checks if item in the storage.
func (d *Disk) IsItemExist(hash string) (bool, string) {
	for _, ext := range d.imgTypes {
		fileName := d.buildFileName(hash, ext)
		if _, err := os.Stat(fileName); err == nil {
			return true, fileName
		}
	}

	return false, ""
}

func (d *Disk) buildFileName(hash, ext string) string {
	sb := strings.Builder{}
	sb.WriteString(hash)
	sb.WriteByte('.')
	sb.WriteString(ext)

	return filepath.Join(d.storPath, sb.String())
}
