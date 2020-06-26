package api

import (
	"fmt"
	"gin-admin/middleware"
	"gin-admin/service"
	"gin-admin/utils"
	"gin-admin/utils/response"
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
		response.FailMessage(err.Error(), c)
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":  0,
			"Menus": menus,
		})
	}
}

// @Tags menu
// @Summary 分页获取基础menu列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.PageInfo true "分页获取基础menu列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /menu/getMenuList [post]
func GetMenuList(c *gin.Context) {
	var pageInfo utils.PageInfo
	var err error
	_ = c.ShouldBindJSON(&pageInfo)
	utils.Initvalidate()
	err = utils.Validate.Struct(pageInfo)
	if err != nil {
		utils.HandleError(err, c)
		return
	}

	err, menuList, total := service.GetMenuList()
	if err != nil {
		response.FailMessage(fmt.Sprintf("获取数据失败，%v", err), c)
		return
	}
	data := utils.PageResult{
		List:     menuList,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}
	response.OkData(data, c)

}
