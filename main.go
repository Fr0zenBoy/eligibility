package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Fr0zenBoy/authoraizer/server"
)

func main() {
	router := gin.New()
	router.POST("/api/authoraizer", server.AuthoraizerHandler)
	router.Run("localhost:8080")
}
