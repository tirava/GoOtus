// Package storages implements any storages.
package storages

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // postgres driver
	"gitlab.com/tirava/shop/internal/models"
)

// GormDB base struct.
type GormDB struct {
	DB *gorm.DB
}

// NewGormDB returns new GORM db.
func NewGormDB(dia, dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(dia, dsn)
	if err != nil {
		return nil, err
	}

	db.AutoMigrate(&models.User{})

	return db, nil
}
