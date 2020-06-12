package router

import (
	. "gin-admin/api"
	"gin-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitUserRouter(Router *gin.RouterGroup) {
	UserRouter := Router.Group("user").Use(middleware.JWTAuth())
	{
		UserRouter.GET("/info", User)
		UserRouter.GET("/list", ListUser)
		UserRouter.POST("/add", AddUser)
		UserRouter.POST("/edit", EditUser)
		UserRouter.POST("/delete", DeleteUser)
	}
}
