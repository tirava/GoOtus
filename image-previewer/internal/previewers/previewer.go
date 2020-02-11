// Package previewers implements image interface.
package previewers

import (
	"fmt"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/preview"
	"gitlab.com/tirava/image-previewer/internal/models"
	nfntcrop "gitlab.com/tirava/image-previewer/internal/previewers/nfnt_crop"
	"gitlab.com/tirava/image-previewer/internal/previewers/xdraw"
)

// NewPreviewer returns previewer implementer.
func NewPreviewer(implementer string) (preview.Previewer, error) {
	switch implementer {
	case models.NfntCrop:
		return nfntcrop.NewNfNtCrop()
	case models.XDraw:
		return xdraw.NewXDraw()
	}

	return nil, fmt.Errorf("incorrect previewer implementer name: %s", implementer)
}
