package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Creates a Gin router with default middleware (logger and recovery)

	r.GET("/", handleRequest)

	r.Run()
}

func handleRequest(c *gin.Context) {
	c.JSON(200, gin.H{
		"key": "value",
	})
}
