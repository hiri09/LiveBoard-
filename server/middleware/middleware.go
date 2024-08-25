package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type usercred struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func LoginMiddleWare(c *gin.Context) {

	var userDetails usercred
	err := c.ShouldBindBodyWithJSON(&userDetails)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username or Password is missing",
		})
		return
	}

	if userDetails.Username == "" || userDetails.Password == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Username or Password is missing",
		})
		return
	}

	c.Next()
}
