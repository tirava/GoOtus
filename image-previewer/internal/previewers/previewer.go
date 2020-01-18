/*
 * Project: Image Previewer
 * Created on 17.01.2020 11:22
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package previewers implements image interface.
package previewers

import (
	"errors"

	nfntcrop "gitlab.com/tirava/image-previewer/internal/previewers/nfnt_crop"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/preview"
)

// NewPreviewer returns previewer implementer.
func NewPreviewer(implementer string) (preview.Previewer, error) {
	switch implementer {
	case "nfnt_crop":
		return nfntcrop.NewNfNtCrop()
	case "other":
	}

	return nil, errors.New("incorrect implementer name")
}
