package main

import (
	"gin-admin/config"
	"gin-admin/db"
	"gin-admin/initialize"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	db.Initialize()

}

// @title Swagger Example API
// @version 0.0.1
// @description This is a sample Server pets
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @BasePath /
func main() {
	gin.SetMode(gin.DebugMode)
	log.Println(config.ApplicationConfig.Port)
	r := initialize.Routers()
	defer db.Orm.Close()
	if err := r.Run(config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port); err != nil {
		log.Fatal(err)
	}
}
