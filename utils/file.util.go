package utils

import (
	"UserPhoto-API/constanta"
	"fmt"
	"log"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func FileValidation(fileHeader *multipart.FileHeader, fileType []string) bool {

	contentType := fileHeader.Header.Get("Content-Type")
	log.Println("conten-type", contentType)
	result := false

	for _, typeFile := range fileType {
		if contentType == typeFile {
			result = true
			break
		}
	}

	return result

}

func FileValidationByExtension(fileHeader *multipart.FileHeader, fileExtension []string) bool {

	extension := filepath.Ext(fileHeader.Filename)
	log.Println("conten-type", extension)
	result := false

	for _, typeFile := range fileExtension {
		if extension == typeFile {
			result = true
			break
		}
	}

	return result

}

func SaveFile(ctx *gin.Context, fileHeader *multipart.FileHeader, filename string) bool {

	errUpload := ctx.SaveUploadedFile(fileHeader, fmt.Sprintf("%s%s", constanta.DIR_IMAGE, filename))
	if errUpload != nil {
		log.Println("can't save file")
		return false
	} else {
		return true
	}
}

func RemoveFile(filePath string) error {
	err := os.Remove(filePath)
	if err != nil {
		log.Println("Failed to remove file")
		return err
	}

	return nil

}
