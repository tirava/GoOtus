// Package hash implements encoding hash interface.
package hash

import (
	"encoding/hex"
	"fmt"
	"hash"
)

// Hash is the base hash type.
type Hash struct {
	hasher hash.Hash
}

// NewHash returns new hash struct.
func NewHash(hasher hash.Hash) (*Hash, error) {
	return &Hash{hasher: hasher}, nil
}

// Encode hashes url to hash.
func (h Hash) Encode(imageURL string) (string, error) {
	h.hasher.Reset()

	_, err := h.hasher.Write([]byte(imageURL))
	if err != nil {
		return "", fmt.Errorf("error write in hasher: %w", err)
	}

	return hex.EncodeToString(h.hasher.Sum(nil)), nil
}
