/*
 * Project: Image Previewer
 * Created on 26.01.2020 15:33
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package encoders implements encoders interface.
package encoders

import (
	"errors"

	"gitlab.com/tirava/image-previewer/internal/encoders/sha256hash"

	"gitlab.com/tirava/image-previewer/internal/encoders/sha1hash"

	"gitlab.com/tirava/image-previewer/internal/encoders/md5hash"

	"gitlab.com/tirava/image-previewer/internal/domain/interfaces/encode"
)

// NewImageURLEncoder returns previewer implementer.
func NewImageURLEncoder(implementer string) (encode.Hasher, error) {
	switch implementer {
	case "md5":
		return md5hash.NewHash()
	case "sha1":
		return sha1hash.NewHash()
	case "sha256":
		return sha256hash.NewHash()
	}

	return nil, errors.New("incorrect url encoder implementer name")
}
