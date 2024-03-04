package boostrap

import (
	"UserPhoto-API/config"
	"UserPhoto-API/config/app_config"
	"UserPhoto-API/config/cors_config"
	"UserPhoto-API/config/log_config"
	"UserPhoto-API/database"
	"log"

	routes "UserPhoto-API/routers"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func BoostrapApp() {
	// LOAD ENV FILE
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	// INIR CONFIG
	config.InitConfig()

	// DATABASE CONNECTION
	database.ConnetDatabase()

	// Move log to folder logs
	log_config.DefaultLogging()
	log.Println("====================================")

	// INIT GIN ENGINE
	app := gin.Default()

	// CORS
	app.Use(cors_config.CorsConfigContrib())
	// app.Use(cors.Config.CorsConfig)

	// INIT ROUTE
	routes.InitRoute(app)

	// RUN APP
	app.Run(app_config.PORT)

}
