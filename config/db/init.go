// Package db : db configuration
package db

import (
	"gorm.io/gorm"

	domain "github.com/aziemp66/go-gorm/internal/domain"
)

// AutoMigrate : Automatically Migrate Database tables
func AutoMigrate(db *gorm.DB) {
	company := &domain.Company{}
	user := &domain.User{}

	err := db.AutoMigrate(company)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(user)
	if err != nil {
		panic(err)
	}
}
