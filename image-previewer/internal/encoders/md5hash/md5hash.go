/*
 * Project: Image Previewer
 * Created on 26.01.2020 15:57
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package md5hash implements md5 hash algorithm.
package md5hash

// nolint md5 leak
import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

// MD5 is the base md5 type.
type MD5 struct {
}

// NewHash returns new md5 struct.
func NewHash() (*MD5, error) {
	return &MD5{}, nil
}

// Encode hashes url to md5.
func (m MD5) Encode(imageURL string) (string, error) {
	// nolint md5 leak
	hasher := md5.New()

	_, err := hasher.Write([]byte(imageURL))
	if err != nil {
		return "", fmt.Errorf("error write md5 hasher: %w", err)
	}

	return hex.EncodeToString(hasher.Sum(nil)), nil
}
