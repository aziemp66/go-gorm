// Package main : main of the program
package main

import (
	"github.com/gin-gonic/gin"

	DBConfig "github.com/aziemp66/go-gorm/config/db"
	DB "github.com/aziemp66/go-gorm/pkg/postgres"
)

func main() {
	router := gin.Default()

	db := DB.NewDB()
	DBConfig.AutoMigrate(db)

	err := router.Run(":5000")
	if err != nil {
		panic(err)
	}
}
