package router

import (
	. "gin-admin/api"
	"gin-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.JWTAuth())
	{
		UserRouter.GET("/info", middleware.JWTAuth(), User)
		UserRouter.GET("/list", ListUser)
		UserRouter.POST("/", AddUser)
		UserRouter.POST("/edit", middleware.JWTAuth(), EditUser)
	}
}
