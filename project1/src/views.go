package main

import (
	"./models"
	"github.com/gin-gonic/gin"
	"log"
)

func changeImgView(request *gin.Context) {
	file, err := request.FormFile("img_file")
	if err != nil {
		log.Println("Error loading file: ", err)
		request.String(errorResponse("Cannot upload file"))
	}

	log.Println("File loaded: ", file.Filename)
	fileName, filePath := generateFilePath()
	if err = request.SaveUploadedFile(file, filePath); err != nil {
		log.Println("Error saving of file: ", err)
		request.String(errorResponse("Cannot save file"))
	}

	size, errorMessage := models.GetNewSize(request.PostForm("width"), request.PostForm("height"))
	if errorMessage != "" {
		request.String(errorResponse(errorMessage))
	}

	format, errorMessage := getFormat(request.PostForm("format"))
	if errorMessage != "" {
		request.String(errorResponse(errorMessage))
	}

	newFilePath, errorMessage := models.Build(fileName, filePath, size, format)
	if errorMessage != "" {
		request.String(errorResponse(errorMessage))
	}

	request.File(newFilePath)
}
