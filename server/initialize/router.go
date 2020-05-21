package initialize

import (
	"gin-admin/middleware"
	"gin-admin/router"

	_ "gin-admin/docs"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func Routers() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Cors())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	ApiGroup := r.Group("")
	router.InitBaseRouter(ApiGroup)
	router.InitUserRouter(ApiGroup)
	router.InitMenuRouter(ApiGroup)
	return r
}
