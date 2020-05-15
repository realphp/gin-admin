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

func main() {
	gin.SetMode(gin.DebugMode)
	log.Println(config.ApplicationConfig.Port)
	r := initialize.Routers()
	defer db.Orm.Close()
	if err := r.Run(config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port); err != nil {
		log.Fatal(err)
	}
}
