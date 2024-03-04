package middleware

import (
	"UserPhoto-API/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware(ctx *gin.Context) {
	bearerToken := ctx.GetHeader("Authorization")

	if !strings.Contains(bearerToken, "Bearer") {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "invalid token",
		})
		return
	}

	token := strings.Replace(bearerToken, "Bearer ", "", -1)
	if token == "" {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated token",
		})
		return
	}
	claimsData, err := utils.DecodedToken(token)
	if err != nil {
		ctx.AbortWithStatusJSON(401, gin.H{
			"message": "unauthenticated claimsData",
			"err":     err,
		})
		return
	}

	ctx.Set("claimsData", claimsData)
	ctx.Set("user_id", claimsData["id"])

	ctx.Next()

}

func TokenMiddleware(ctx *gin.Context) {

}
