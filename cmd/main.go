// Package main : main of the program
package main

import (
	"github.com/gin-gonic/gin"

	_ "github.com/aziemp66/go-gorm/pkg/postgres"
)

func main() {
	router := gin.Default()

	// db := DB.NewDB()

	err := router.Run(":5000")
	if err != nil {
		panic(err)
	}
}
