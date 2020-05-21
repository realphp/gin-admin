package api

import (
	"gin-admin/middleware"
	"gin-admin/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @Tags authorityAndMenu
// @Summary 获取用户动态路由
// @Security ApiKeyAuth
// @Produce  application/json
// @Param data body request.RegisterAndLoginStruct true "可以什么都不填"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"返回成功"}"
// @Router /menu/getMenu [post]
func GetMenu(c *gin.Context) {
	claims, _ := c.Get("claims")
	waitUse := claims.(*middleware.UserClaim)
	err, menus := service.GetMenuTree(waitUse.RoleId)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code": 0,
			"msg":  err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"Menus": menus,
		})
	}
}
