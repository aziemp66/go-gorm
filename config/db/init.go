// Package db : db configuration
package db

import (
	"gorm.io/gorm"

	domain "github.com/aziemp66/go-gorm/internal/domain"
)

// AutoMigrate : Automatically Migrate Database tables
func AutoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&domain.Company{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.User{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.Customer{})
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&domain.CreditCard{})
	if err != nil {
		panic(err)
	}
}
