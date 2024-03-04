package user_controller

import (
	"UserPhoto-API/database"
	"UserPhoto-API/models/full_model"
	"UserPhoto-API/models/request_model"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context) {
	userReq := new(request_model.UserRequest)

	// CHECK INPUT FORM
	errReq := ctx.ShouldBind(&userReq)
	if errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	// CHECK EMAIL MUST UNIQUE
	userEmailExist := new(full_model.User)
	database.DB.Table("user").Where("email = ?", userReq.Email).Find(&userEmailExist)
	fmt.Println(userEmailExist.ID)
	if userEmailExist.ID == nil {
		userEmailExist = nil
	} else if userEmailExist != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "email is already used",
		})
		return
	}

	// CHECK MINLENGTH CHARACTER OF PASSWORD (>=6)
	if len(userReq.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "minimal character of password is 6 required",
		})
		return
	}

	user := new(full_model.User)
	user.Username = &userReq.Username
	user.Email = &userReq.Email
	user.Password = &userReq.Password

	// STORE DATA
	errDb := database.DB.Table("user").Create(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "Internal server error -> cant stored data",
		})
		return
	}

	// SHOW RESPONSE DATA
	getUserDb := database.DB.Table("user").Where("email = ?", user.Email).Find(&user).Error
	if getUserDb != nil {
		ctx.JSON(500, gin.H{
			"message": "data stored successfully",
			"massage": "Internal server error -> cant get the data",
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"message": "data stored successfully",
			"data":    user,
		})
	}

}

func UpdateUser(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userID => ", userId)

	id := ctx.Param("userId")

	if strconv.FormatFloat(userId, 'f', -1, 64) != id {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "access rejected : token didnt valid",
		})
		return
	}

	user := new(full_model.User)
	userReq := new(request_model.UserRequest)
	userEmailExist := new(full_model.User)

	// CEK ERROR USER INPUT REQUEST
	errReq := ctx.ShouldBind(&userReq)
	if errReq != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	// CHECK QUERRY USER ID EXIST
	errDb := database.DB.Table("user").Where("id = ?", id).Find(&user).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	// CHECK USER ID EXIST
	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"message": "data not found",
		})
		return
	}

	// CHECK EMAIL EXIXST (EMAIL -> UNIQUE)
	errUserEmailExist := database.DB.Table("user").Where("email = ?", userReq.Email).Find(&userEmailExist).Error
	if errUserEmailExist != nil {
		ctx.JSON(500, gin.H{
			"message": "internal server error",
		})
		return
	}
	if userEmailExist.Email != nil && *user.ID != *userEmailExist.ID {
		ctx.JSON(400, gin.H{
			"message": "Email Already Used",
		})
		return
	}

	// CHECK MINLENGTH CHARACTER OF PASSWORD (>=6)
	if len(userReq.Password) < 6 {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "minimal character of password is 6 required",
		})
		return
	}

	currentTime := time.Now()

	user.Username = &userReq.Username
	user.Email = &userReq.Email
	user.Password = &userReq.Password
	user.UpdateAt = &currentTime

	// UPDATE DATA USER IN DATABASE
	errUpdate := database.DB.Table("user").Where("id = ?", id).Updates(&user).Error
	if errUpdate != nil {
		ctx.JSON(500, gin.H{
			"message": "cant update data",
		})
		return
	}

	// SHOW RESPONSE DATA
	getUserDb := database.DB.Table("user").Where("id = ?", user.ID).Find(&user).Error
	if getUserDb != nil {
		ctx.JSON(400, gin.H{
			"message": "data stored successfully",
			"massage": "Internal server error -> cant get the data",
		})
		return
	} else {
		ctx.JSON(200, gin.H{
			"message": "data stored successfully",
			"data":    user,
		})
	}

}

func DeleteUser(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userID => ", userId)

	id := ctx.Param("userId")

	if strconv.FormatFloat(userId, 'f', -1, 64) != id {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "access rejected : token didnt valid",
		})
		return
	}

	user := new(full_model.User)

	// CHECK DATA IS EXIST
	errCheck := database.DB.Table("user").Where("id = ?", id).Find(&user).Error
	if errCheck != nil {
		ctx.JSON(500, gin.H{
			"meesage": "internal server error",
			"error":   errCheck.Error(),
		})
		return
	}
	if user.ID == nil {
		ctx.JSON(404, gin.H{
			"meesage": "data not found",
		})
		return
	}

	// DELETE DATA FROM DATABASE
	errDb := database.DB.Table("user").Unscoped().Where("id = ?", id).Delete(&full_model.User{}).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"meesage": "internal server error",
			"error":   errDb.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"meesage": "data deleted successfully",
	})

}
