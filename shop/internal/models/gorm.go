// Package models implements external entity models.
package models

import (
	"github.com/jinzhu/gorm"
	"gitlab.com/tirava/shop/internal/domain/entities"
)

// User models struct.
type User struct {
	gorm.Model
	entities.User
}
