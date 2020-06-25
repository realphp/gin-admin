package api

import (
	"gin-admin/model"
	"gin-admin/service"
	"gin-admin/utils"
	"gin-admin/utils/response"

	"github.com/gin-gonic/gin"
)

func ListRole(c *gin.Context) {
	var pageInfo utils.PageInfo
	err := c.ShouldBindQuery(&pageInfo)
	if err != nil {
		response.FailMessage(err.Error(), c)
	}
	err, list, total := service.ListRole(pageInfo)
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

func AddRole(c *gin.Context) {
	var role model.Role
	var err error
	_ = c.ShouldBindJSON(&role)
	utils.Initvalidate()
	err = utils.Validate.Struct(role)
	if err != nil {
		utils.HandleError(err, c)
		return
	}
	err, role = service.AddRole(role)
	if err != nil {
		response.FailMessage(err.Error(), c)
		return
	}
	response.Ok(c)

}

func EditRole(c *gin.Context) {
	var role model.Role
	var err error
	err = c.ShouldBindJSON(&role)
	if role.Id == 0 {
		response.FailMessage("请求参数错误", c)
		return
	}
	err = service.UpdateRole(&role)
	if err != nil {
		response.FailMessage(err.Error(), c)
		return
	}
	response.Ok(c)

}

func DeleteRole(c *gin.Context) {
	var role model.Role
	_ = c.ShouldBindJSON(&role)
	if role.Id == 0 {
		response.FailMessage("请求参数错误", c)
		return
	}
	err := service.DeleteRole(&role)
	if err != nil {
		response.FailMessage(err.Error(), c)
		return
	}
	response.Ok(c)
}
