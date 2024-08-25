package router

import (
	"go-practice/middleware"

	"github.com/gin-gonic/gin"
)

func RouteIncomingurl(r *gin.Engine) {
	r.POST("/SingIn", middleware.LoginMiddleWare)
}
