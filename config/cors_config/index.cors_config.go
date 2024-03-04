package cors_config

import (
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var origins = []string{
	"https://domain-saya.com",
	"https://sub.domain-saya.com",
}

func CorsConfigContrib() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = origins

	return cors.New(config)

}

func CorsConfig(c *gin.Context) {

	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "true")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "Content-Type, Content-length")
	c.Writer.Header().Set("Access-Control-Allow-Origin", "POST, GET, PATCH, DELETE, PUT, OPTION")

	if c.Request.Method == "OPTION" {
		c.AbortWithStatus(http.StatusNoContent)
		return
	}

	c.Next()

}
