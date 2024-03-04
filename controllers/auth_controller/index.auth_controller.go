package auth_controller

import (
	"UserPhoto-API/database"
	"UserPhoto-API/models/full_model"
	"UserPhoto-API/models/request_model"
	"UserPhoto-API/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

func Login(ctx *gin.Context) {
	loginReq := new(request_model.LoginRequest)

	errReq := ctx.ShouldBind(&loginReq)
	if errReq != nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	// CHECK EMAIL IS EXIST
	user := new(full_model.User)
	errUser := database.DB.Table("user").Where("email = ?", loginReq.Email).Find(&user).Error
	if errUser != nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "email doenst exist",
		})
		return
	}
	if user.Email == nil {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "email doenst exist",
		})
		return
	}

	// CHECK PASSWORD
	if loginReq.Password != *user.Password {
		ctx.AbortWithStatusJSON(404, gin.H{
			"message": "password not valid",
		})
		return
	}

	// GET MAP CLAIM
	claims := jwt.MapClaims{
		"id":  user.ID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	}

	token, errToken := utils.GenerateToken(&claims)
	if errToken != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"message": "failed generated token",
		})
		return
	}

	ctx.JSON(200, gin.H{
		"mmessage": "login succesfully",
		"token":    token,
	})

}
