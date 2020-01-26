/*
 * Project: Image Previewer
 * Created on 26.01.2020 14:53
 * Copyright (c) 2020 - Eugene Klimov
 */

// Package encode implements encode interface.
package encode

// Hasher is the main interface for name hash logic.
type Hasher interface {
	Encode(imageURL string) (string, error)
}
