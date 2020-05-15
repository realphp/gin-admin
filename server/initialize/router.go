package initialize

import (
	"gin-admin/router"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	ApiGroup := r.Group("")
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	return r
}
