package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	err := router.Run(":5000")
	if err != nil {
		panic(err)
	}
}
