// Package encode implements encode interface.
package encode

// Hasher is the main interface for name hash logic.
type Hasher interface {
	Encode(imageURL string) (string, error)
}
