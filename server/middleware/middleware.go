package middleware

import "github.com/gin-gonic/gin"

type usercred struct {
	Username string `json:"username"`
	Password string `json:"Password"`
}

func LoginMiddleWare(c *gin.Context) {

	// var userdata usercred

}
