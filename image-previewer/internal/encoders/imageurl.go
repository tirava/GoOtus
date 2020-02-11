// Package encoders implements encoders interface.
package encoders

// nolint:gosec
import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/encode"
	"gitlab.com/tirava/image-previewer/internal/encoders/hash"
	"gitlab.com/tirava/image-previewer/internal/models"
)

// NewImageURLEncoder returns previewer implementer.
func NewImageURLEncoder(implementer string) (encode.Hasher, error) {
	switch implementer {
	case models.MD5:
		return hash.NewHash(md5.New()) // nolint:gosec
	case models.SHA1:
		return hash.NewHash(sha1.New()) // nolint:gosec
	case models.SHA256:
		return hash.NewHash(sha256.New())
	}

	return nil, fmt.Errorf("incorrect url encoder implementer name: %s", implementer)
}
