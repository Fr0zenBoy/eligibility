package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Fr0zenBoy/authoraizer/routes"
)

func main() {
	router := gin.New()
	router.POST("/api/authoraizer", routes.AuthoraizerHandler)
	router.Run("localhost:8080")
}
