/*
 * Project: Image Previewer
 * Created on 17.01.2020 11:22
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package previewers implements image interface.
package previewers

import (
	"errors"

	"gitlab.com/tirava/image-previewer/internal/previewers/xdraw"

	nfntcrop "gitlab.com/tirava/image-previewer/internal/previewers/nfnt_crop"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/preview"
)

// NewPreviewer returns previewer implementer.
func NewPreviewer(implementer string) (preview.Previewer, error) {
	switch implementer {
	case "nfnt_crop":
		return nfntcrop.NewNfNtCrop()
	case "xdraw":
		return xdraw.NewXDraw()
	}

	return nil, errors.New("incorrect preview implementer name")
}
