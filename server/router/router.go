package router

import (
	. "gin-admin/api"
	"gin-admin/middleware"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Static("/static", "./static")
	r.POST("/login", Login)
	r.GET("/user", middleware.JWTAuth(), User)
	r.GET("/user/list", ListUser)
	r.POST("/user", AddUser)
	r.POST("/user/edit", middleware.JWTAuth(), EditUser)
	return r
}
