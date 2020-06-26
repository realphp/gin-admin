package api

import (
	"fmt"
	"gin-admin/middleware"
	"gin-admin/model"
	"gin-admin/service"
	"gin-admin/utils"
	"gin-admin/utils/response"

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
	response.OkData(waitUse, c)
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
		response.FailMessage(err.Error(), c)
		return
	}
	response.Ok(c)

}

func EditUser(c *gin.Context) {
	var user model.User
	var err error
	err = c.ShouldBindJSON(&user)
	if user.Id == 0 {
		response.FailMessage("请求参数错误", c)
		return
	}
	err = service.UpdateUser(&user)
	if err != nil {
		response.FailMessage(err.Error(), c)
		return
	}
	response.Ok(c)

}

func DeleteUser(c *gin.Context) {
	var user model.User
	_ = c.ShouldBindJSON(&user)
	if user.Id == 0 {
		response.FailMessage("请求参数错误", c)
		return
	}
	err := service.DeleteUser(&user)
	if err != nil {
		response.FailMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}

func ListUser(c *gin.Context) {
	var pageInfo utils.PageInfo

	err := c.ShouldBindQuery(&pageInfo)
	fmt.Println(err, pageInfo)
	err, list, total := service.ListUser(pageInfo)
	if err != nil {
		response.FailMessage(err.Error(), c)
		return
	}
	data := utils.PageResult{
		List:     list,
		Total:    total,
		Page:     pageInfo.Page,
		PageSize: pageInfo.PageSize,
	}
	response.OkData(data, c)
}
