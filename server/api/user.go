package api

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
)



func User(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(jwt.Claims)
	fmt.Println(waitUse)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "asdasd",
	})
}
