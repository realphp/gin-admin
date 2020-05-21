package router

import (
	"gin-admin/api"
	"gin-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitMenuRouter(Router *gin.RouterGroup) (R gin.IRoutes) {
	MenuRouter := Router.Group("menu").Use(middleware.JWTAuth())
	{
		MenuRouter.POST("list", api.GetMenu) //获取菜单树
	}
	return MenuRouter
}
