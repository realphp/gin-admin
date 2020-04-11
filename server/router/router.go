package router

import (
	. "gin-admin/api"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Static("/static", "./static")
	r.GET("/login", Login)
	return r
}
