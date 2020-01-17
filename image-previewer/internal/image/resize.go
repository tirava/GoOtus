/*
 * Project: Image Previewer
 * Created on 17.01.2020 11:22
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package image implements image interface.
package image

import (
	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/image"
	"gitlab.com/tirava/image-previewer/internal/image/nfnt"
)

// NewResizer returns resizer implementer.
func NewResizer(implementer string) (image.Resizer, error) {
	switch implementer {
	case "nfnt":
		return nfnt.NewNfNt()
	case "other":
	}

	return nil, nil
}
