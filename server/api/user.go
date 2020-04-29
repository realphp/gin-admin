package api

import (
	"fmt"
	"gin-admin/initialize"
	"gin-admin/middleware"
	"gin-admin/model"
	"gin-admin/service"
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

func AddUser(c *gin.Context) {
	var user model.User
	var err error
	err = c.ShouldBindJSON(&user)
	err = initialize.Validate.Struct(user)
	if err != nil {
		initialize.HandleError(err, c)
		return
	}
	err = service.AddUser(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "创建成功",
	})

}
