package config

import (
	"github.com/spf13/viper"
	"log"
)

var cfgDatabase *viper.Viper
var cfgApplication *viper.Viper
var cfgRedisConn *viper.Viper

func init() {
	viper.SetConfigName("env") // name of config file (without extension)
	viper.AddConfigPath("./") // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {
		log.Println(err) // Handle errors reading the config file
	}

	cfgDatabase = viper.Sub("settings.database")
	if cfgDatabase == nil {
		panic("config not found settings.database")
	}
	DatabaseConfig = InitDatabase(cfgDatabase)

	cfgApplication = viper.Sub("settings.application")
	if cfgApplication == nil {
		panic("config not found settings.application")
	}
	ApplicationConfig = InitApplication(cfgApplication)

	cfgRedisConn = viper.Sub("settings.redis")
	if cfgRedisConn == nil {
		panic("config not found settings.redis")
	}
	RedisConnConfig = InitRedisConn(cfgRedisConn)

}

func SetApplicationIsInit() {
	viper.AddConfigPath("./config")
	viper.Set("settings.application.isInit", false)
	viper.WriteConfig()
}
