package router

import (
	. "gin-admin/api"

	"github.com/gin-gonic/gin"
)

func InitBaseRouter(Router *gin.RouterGroup) {
	Router.Static("/static", "./static")
	Router.POST("/login", Login)
}
