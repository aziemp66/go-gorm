package main

import (
	"log"

	"github.com/gin-gonic/gin"

	DB "github.com/aziemp66/go-gorm/pkg/postgres"
)

func main() {
	router := gin.Default()

	db := DB.NewDB()

	log.Println(db)

	err := router.Run(":5000")
	if err != nil {
		panic(err)
	}
}
