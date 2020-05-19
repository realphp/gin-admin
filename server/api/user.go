package api

import (
	"fmt"
	"gin-admin/middleware"
	"gin-admin/model"
	"gin-admin/service"
	"gin-admin/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags User
// @Summary 用户信息
// accept	json
// produce	json
// @Success 200 {string} string	"ok"
// @Router /user/ [get]
func User(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.UserClaim)
	fmt.Println(waitUse.Username)
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"data": waitUse,
	})
}

func AddUser(c *gin.Context) {
	var user model.User
	var err error
	err = c.ShouldBindJSON(&user)
	utils.Initvalidate()
	err = utils.Validate.Struct(user)
	if err != nil {
		utils.HandleError(err, c)
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

func EditUser(c *gin.Context) {
	var user model.User
	var err error
	claims, _ := c.Get("claims")
	User := claims.(*middleware.UserClaim)
	err = c.ShouldBindJSON(&user)
	err = service.UpdateUser(User.Id, &user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "修改成功",
	})

}

func ListUser(c *gin.Context) {
	var pageInfo utils.PageInfo

	err := c.ShouldBindQuery(&pageInfo)
	fmt.Println(err, pageInfo)
	err, list, total := service.ListUser(pageInfo)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	}
	data := utils.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"data": data,
	})
}
