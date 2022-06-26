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
	currFormat := getCurrImgFormat(file.Filename)
	fileName, filePath := generateFilePath(currFormat)

	if err = request.SaveUploadedFile(file, filePath); err != nil {
		log.Println("Error saving of file: ", err)
		request.String(errorResponse("Cannot save file"))
	}

	size, errorMessage := models.GetNewSize(request.PostForm("width"), request.PostForm("height"))
	if errorMessage != "" {
		request.String(errorResponse(errorMessage))
	}

	format := models.Format{Old: currFormat, New: request.PostForm("format")}
	newFilePath, errorMessage := models.Build(fileName, filePath, size, format)
	if errorMessage != "" {
		request.String(errorResponse(errorMessage))
	}

	request.File(newFilePath)
}
