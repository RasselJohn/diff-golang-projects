package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sunshineplan/imgconv"
	"log"
	"math/rand"
	"net/http"
)

const ImgLoadFolder = "./static/loaded/"
const ImgResultFolder = "./static/result/"

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 20480 // 20 MiB
	router.POST("/change-image", uploadImg)
	router.Run() // listen and serve on 0.0.0.0:8080
}

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

	src, err := imgconv.Open(filePath)
	if err != nil {
		log.Println("Error opening file: ", err)
		request.String(errorResponse("Cannot convert file"))
	}

	err = imgconv.Save(
		fmt.Sprintf("%v%v.png", ImgResultFolder, randomFileName),
		src,
		imgconv.FormatOption{Format: imgconv.PNG},
	)
	if err != nil {
		log.Println("Error converting of file: ", err)
		request.String(errorResponse("Cannot save the converted file"))
	}
	request.String(http.StatusOK, fmt.Sprintf("File converted!"))
}

func errorResponse(reason string) (statusCode int, message string) {
	return http.StatusBadRequest, fmt.Sprintf("Error happend. Reason: '%v'. Please, repeat later.", message)
}
