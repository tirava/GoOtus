/*
 * Project: Image Previewer
 * Created on 26.01.2020 15:57
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package sha256hash implements md5 hash algorithm.
package sha256hash

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

// SHA256 is the base sha1 type.
type SHA256 struct {
}

// NewHash returns new sha256 struct.
func NewHash() (*SHA256, error) {
	return &SHA256{}, nil
}

// Encode hashes url to sha1.
func (s256 SHA256) Encode(imageURL string) (string, error) {
	hasher := sha256.New()

	_, err := hasher.Write([]byte(imageURL))
	if err != nil {
		return "", fmt.Errorf("error write sha256 hasher: %w", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
