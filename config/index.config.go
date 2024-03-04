package config

import (
	"UserPhoto-API/config/app_config"
	"UserPhoto-API/config/db_config"
)

func InitConfig() {
	app_config.InitAppConfig()
	db_config.InitDatabaseConfig()
	// log_config.DefaultLogging()
}
