package log_config

import (
	"io"
	"log"
	"os"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

var defaultLogFilePath = "logs/file/gin.log"

func createFolderIfNotExist(path string) {
	dir := filepath.Dir(path)

	_, err := os.Stat(dir)
	if os.IsNotExist(err) {
		log.Println("Create", dir, "directory")
		os.MkdirAll(dir, 0644)
		if err != nil {
			log.Println("Failed ro create", dir)
		} else {
			log.Printf(dir, "directory created")
		}
	}

}

func openOrCreatedLogFile(path string) (*os.File, error) {
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)

	if err != nil {
		var errCreateFile error
		logFile, errCreateFile = os.Create(path)
		if errCreateFile != nil {
			log.Println("cant create log file", errCreateFile)
		}
	}

	return logFile, nil

}

func DefaultLogging() {
	gin.DisableConsoleColor()

	createFolderIfNotExist(defaultLogFilePath)
	f, _ := openOrCreatedLogFile(defaultLogFilePath)
	gin.DefaultWriter = io.MultiWriter(f)
	log.SetOutput(gin.DefaultWriter)

}
