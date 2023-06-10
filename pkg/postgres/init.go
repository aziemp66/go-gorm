// Package db : Database Config
package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	config "github.com/aziemp66/go-gorm/config/db"
)

// NewDB : Return New Database Config
func NewDB() *gorm.DB {
	dsn := "host=localhost user=aziemp66 password=azie122333 dbname=go_gorm port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	log.Println("Database Connected")

	config.AutoMigrate(db)

	return db
}
