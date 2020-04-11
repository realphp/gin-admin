package main

import (
	"gin-admin/config"
	"gin-admin/db"
	"gin-admin/router"
	"github.com/gin-gonic/gin"
	"log"
)

func init() {
	db.Initialize()
}

func main() {
	gin.SetMode(gin.DebugMode)
	log.Println(config.ApplicationConfig.Port)
	r := router.InitRouter()
	defer db.Orm.Close()
	if err := r.Run(config.ApplicationConfig.Host + ":" + config.ApplicationConfig.Port); err != nil {
		log.Fatal(err)
	}
}
