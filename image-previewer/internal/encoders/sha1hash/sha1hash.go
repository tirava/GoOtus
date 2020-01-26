/*
 * Project: Image Previewer
 * Created on 26.01.2020 15:57
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package sha1hash implements md5 hash algorithm.
package sha1hash

// nolint sha1 leak
import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// SHA1 is the base sha1 type.
type SHA1 struct {
}

// NewHash returns new sha1 struct.
func NewHash() (*SHA1, error) {
	return &SHA1{}, nil
}

// Encode hashes url to sha1.
func (s1 SHA1) Encode(imageURL string) (string, error) {
	// nolint sha1 leak
	hasher := sha1.New()

	_, err := hasher.Write([]byte(imageURL))
	if err != nil {
		return "", fmt.Errorf("error write sha1 hasher: %w", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
