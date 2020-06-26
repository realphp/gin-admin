package db

import (
	"fmt"
	"gin-admin/config"
	"log"
	"strconv"

	_ "github.com/go-sql-driver/mysql" //加载mysql
	"github.com/jinzhu/gorm"
)

var Orm *gorm.DB

func Initialize() {
	dbType := config.DatabaseConfig.Dbtype
	host := config.DatabaseConfig.Host
	port := strconv.Itoa(config.DatabaseConfig.Port)
	database := config.DatabaseConfig.Database
	username := config.DatabaseConfig.Username
	password := config.DatabaseConfig.Password

	if dbType != "mysql" {
		fmt.Println("db type unknow")
	}
	var err error
	Orm, err = gorm.Open("mysql", username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8&parseTime=True&loc=Local")
	Orm.LogMode(true)
	if err != nil {
		log.Fatalln("%s connect error %v", dbType, err)
	}
	if Orm.Error != nil {
		log.Fatalln("database error %v", Orm.Error)
	}

}
