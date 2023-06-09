// Package db : db configuration
package db

import (
	"gorm.io/gorm"

	"github.com/aziemp66/go-gorm/internal/repository"
)

// AutoMigrate : Automatically Migrate Database tables
func AutoMigrate(db *gorm.DB) {
	company := repository.Company{}
	user := repository.User{}

	err := db.AutoMigrate(company)
	if err != nil {
		panic(err)
	}
	err = db.AutoMigrate(user)
	if err != nil {
		panic(err)
	}
}
