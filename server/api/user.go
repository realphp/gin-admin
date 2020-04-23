package api

import (
	"fmt"
	"gin-admin/middleware"
	"net/http"

	"github.com/gin-gonic/gin"
)

func User(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.UserClaim)
	fmt.Println(waitUse.Username)
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "asdasd",
	})
}
