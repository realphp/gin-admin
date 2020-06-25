package router

import (
	. "gin-admin/api"
	"gin-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRoleRouter(Router *gin.RouterGroup) {
	RoleRouter := Router.Group("role").Use(middleware.JWTAuth())
	{
		RoleRouter.GET("/list", ListRole)
		RoleRouter.POST("/add", AddRole)
		RoleRouter.POST("/edit", EditRole)
		RoleRouter.POST("/delete", DeleteRole)
	}
}
