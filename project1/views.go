package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
)

func uploadImg(request *gin.Context) {
	file, _ := request.FormFile("img_file")
	log.Println("File loaded: ", file.Filename)

	randByteData := make([]byte, 16)
	rand.Read(randByteData)
	randomFileName := hex.EncodeToString(randByteData)
	filePath := fmt.Sprintf("%v%v.jpg", ImgLoadFolder, randomFileName)

	err := request.SaveUploadedFile(file, filePath)
	if err != nil {
		log.Println("Error saving of file: ", err)
		request.String(errorResponse("Cannot save file"))
	}

	_, err := Build(
		randomFileName,
		filePath,
		request.PostForm("resize"),
		request.PostForm("watermark"),
		request.PostForm("convert_type"),
	)
	if err != nil {
		request.String(errorResponse("Cannot convert file"))
	}
	request.String(http.StatusOK, fmt.Sprintf("File converted!"))
}

func errorResponse(reason string) (statusCode int, message string) {
	return http.StatusBadRequest, fmt.Sprintf("Error happend. Reason: '%v'. Please, repeat later.", message)
}
