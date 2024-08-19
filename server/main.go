package main

import (
	"go-practice/router"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default() // Creates a Gin router with default middleware (logger and recovery)
	router.RouteIncomingurl(r)
	r.Run()
}
