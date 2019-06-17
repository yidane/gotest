package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	router := gin.Default()
	err := router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
