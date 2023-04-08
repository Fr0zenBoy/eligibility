package main

import (
	"github.com/gin-gonic/gin"
	"github.com/Fr0zenBoy/eligibility/pkg/routes"
)

func main(){
	router := gin.New()
	router.POST("/api/eligible", routes.EligiableHandler)
	router.Run("localhost:8080")
}
