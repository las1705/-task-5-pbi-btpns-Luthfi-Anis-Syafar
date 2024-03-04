package photo_controller

import (
	"UserPhoto-API/constanta"
	"UserPhoto-API/database"
	"UserPhoto-API/models/full_model"
	"UserPhoto-API/models/request_model"
	"UserPhoto-API/utils"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func HandlerUploadPhoto(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userID => ", userId)

	// CHECK FILE IS FILL
	fileHeader, _ := ctx.FormFile("photo")
	if fileHeader == nil {
		ctx.AbortWithStatusJSON(400, gin.H{
			"messaage": "file is required",
		})
		return
	}

	// CHECK INPUT FORM
	photoReq := new(request_model.PhotoRequest)
	errReq := ctx.ShouldBind(&photoReq)
	if errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	// CHECK FILE TYPE OF EXTESION (IMAGE)
	fileExtension := []string{".jpg", ".jpeg", ".png"}
	isFileValidated := utils.FileValidationByExtension(fileHeader, fileExtension)
	if !isFileValidated {
		ctx.AbortWithStatusJSON(400, gin.H{
			"message": "file is not allowed",
		})
		return
	}

	// NAMING IMAGE FILE
	currentTime := time.Now().UTC().Format("20061206")
	getUserId := strconv.FormatFloat(userId, 'f', -1, 64)
	extensionFile := filepath.Ext(fileHeader.Filename)
	filename := fmt.Sprintf("%s-%s-%s%s", currentTime, getUserId, photoReq.Title, extensionFile)

	// SAVE FILE IN FOLDER PUBLIC
	isSaved := utils.SaveFile(ctx, fileHeader, filename)
	if !isSaved {
		ctx.JSON(500, gin.H{
			"message": "internal server error, cant save file",
		})
		return
	}

	photo := new(full_model.Photo)
	url_for_photo := constanta.DIR_IMAGE + filename
	userId_for_photo := int(userId)
	photo.Title = &photoReq.Title
	photo.Caption = &photoReq.Caption
	photo.PhotoUrl = &url_for_photo
	photo.UserId = &userId_for_photo

	// CHECK DATABASE QUERRY CREATE
	errDb := database.DB.Table("photo").Create(&photo).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "cant stored data",
			"but":     "success upload file",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "file uploaded and data stored succesfully",
		})
	}

}

func HandlerEditPhoto(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userID => ", userId)
	intUserId := int(userId)

	photoId := ctx.Param("photoId")

	// CHECK INPUT FORM
	photoReq := new(request_model.PhotoRequest)
	errReq := ctx.ShouldBind(&photoReq)
	if errReq != nil {
		ctx.JSON(400, gin.H{
			"message": errReq.Error(),
		})
		return
	}

	// CHECK PHOTO EXIST IN DATABASE
	photo := new(full_model.Photo)
	photoDb := database.DB.Table("photo").Where("id = ?", photoId).Find(&photo).Error
	if photoDb != nil {
		ctx.JSON(400, gin.H{
			"message": "photo not found in database",
		})
		return
	}

	// CHECK ACCESS USER ID PHOTO
	intPhotoUserId := photo.UserId
	if *intPhotoUserId != intUserId {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "access rejected : token user didnt valid to this photo",
		})
		return
	}

	getUserId := strconv.FormatFloat(userId, 'f', -1, 64)
	oldFilePath := photo.PhotoUrl
	extensionOldFile := filepath.Ext(*oldFilePath)
	currentTime := time.Now().UTC().Format("20061206")
	newFilename := fmt.Sprintf("%s-%s-%s%s", currentTime, getUserId, photoReq.Title, extensionOldFile)
	newFilePath := constanta.DIR_IMAGE + newFilename

	// CHECK FILE IS FILL OR NOT
	fileHeader, _ := ctx.FormFile("photo")
	if fileHeader == nil { // IF FORM PHOTO IS EMPTY -> JUST RENAME THE OLD FILE
		// RENAME THE OLD FILE

		renamePhoto := os.Rename(*oldFilePath, newFilePath)
		if renamePhoto != nil {
			ctx.JSON(500, gin.H{
				"message": "failed change name photo",
				"erro":    renamePhoto,
			})
			return
		}

	} else { // IF FORM PHOTO IS FILL -> UPLOAD NEW FILE & DELETE OLD ONE
		// CHECK FILE TYPE OF EXTESION (IMAGE)
		fileExtension := []string{".jpg", ".jpeg", ".png"}
		isFileValidated := utils.FileValidationByExtension(fileHeader, fileExtension)
		if !isFileValidated {
			ctx.AbortWithStatusJSON(400, gin.H{
				"message": "photo is not allowed",
			})
			return
		}

		// SAVE NEW FILE IN FOLDER PUBLIC
		isSaved := utils.SaveFile(ctx, fileHeader, newFilename)
		if !isSaved {
			ctx.JSON(500, gin.H{
				"message": "internal server error, cant save photo",
			})
			return
		}

		// REMOVE OLD FILE
		err := utils.RemoveFile(*oldFilePath)
		if err != nil {
			ctx.JSON(500, gin.H{
				"message": err.Error(),
			})
			return
		}

	}

	photo.Title = &photoReq.Title
	photo.Caption = &photoReq.Caption
	photo.PhotoUrl = &newFilePath

	// CHECK DATABASE QUERRY CREATE
	errDb := database.DB.Table("photo").Where("id = ?", photoId).Updates(&photo).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"message": "cant update data",
			"but":     "success update photo",
		})
	} else {
		ctx.JSON(200, gin.H{
			"message": "photo updated and database updated succesfully",
		})
	}

}

func HandlerRemovePhoto(ctx *gin.Context) {
	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userID => ", userId)
	intUserId := int(userId)

	photoId := ctx.Param("photoId")

	// CHECK PHOTO EXIST IN DATABASE
	photo := new(full_model.Photo)
	photoDb := database.DB.Table("photo").Where("id = ?", photoId).Find(&photo).Error
	if photoDb != nil {
		ctx.JSON(400, gin.H{
			"message": "photo not found in database",
		})
		return
	}

	// CHECK ACCESS USER ID PHOTO
	intPhotoUserId := photo.UserId
	if *intPhotoUserId != intUserId {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "access rejected : token user didnt valid to this photo",
		})
		return
	}

	oldFilePath := photo.PhotoUrl

	// DELETE DATA FROM DATABASE
	errDb := database.DB.Table("photo").Unscoped().Where("id = ?", photoId).Delete(&full_model.Photo{}).Error
	if errDb != nil {
		ctx.JSON(500, gin.H{
			"meesage": "internal server error : cant remove photo from database",
			"error":   errDb.Error(),
		})
		return
	}

	err := utils.RemoveFile(*oldFilePath)
	if err != nil {
		ctx.JSON(500, gin.H{
			"message": "cant remove file from server",
			"error":   err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "file removed successfully",
	})

}

// ENDPOINT GET PHOTO : GET ALL PHOTO IN ONE USER ID
func GetPhoto(ctx *gin.Context) {

	userId := ctx.MustGet("user_id").(float64)
	fmt.Println("userID => ", userId)
	intUserId := int(userId)

	photos := new([]full_model.Photo)
	err := database.DB.Table("photo").Where("user_id = ?", intUserId).Find(&photos).Error
	if err != nil {
		ctx.AbortWithStatusJSON(500, gin.H{
			"massage": "bad Request",
		})
	}

	ctx.JSON(200, gin.H{
		"DATA": photos,
	})

}
