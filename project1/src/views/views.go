package views

import (
	"img_convert/consts"
	"img_convert/utils"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func ChangeImgView(request *gin.Context) {
	form, err := request.MultipartForm()
	if err != nil {
		log.Println("Error loading form: ", err)
		request.JSON(http.StatusBadRequest, gin.H{"errors": []string{"Unknown form error."}})
		return
	}

	files := form.File["img_files[]"]
	if len(files) == 0 {
		request.JSON(http.StatusBadRequest, gin.H{"errors": []string{"File field is empty."}})
		return
	}

	fileNamesChannel := make(chan string)
	errorChannel := make(chan string)
	for _, file := range files {
		go utils.FileHandler(request, file, fileNamesChannel, errorChannel)
	}

	var fileNames []string
	var errorMessages []string
	for i := 0; i < len(files); i++ {
		fileName := <-fileNamesChannel
		if fileName != "" {
			fileNames = append(fileNames, fileName)
		}

		errorMessage := <-errorChannel
		if errorMessage != "" {
			errorMessages = append(errorMessages, errorMessage)
		}
	}

	if len(fileNames) == 0 {
		request.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages})
		return
	}

	zipFileName, zipFilePath := utils.GenerateFilePath(consts.ImgResultFolder, "zip")
	if err := utils.CompressToZip(zipFilePath, fileNames); err != nil {
		request.JSON(http.StatusBadRequest, gin.H{"errors": "Cannot compress to zip."})
		return
	}

	request.JSON(http.StatusOK, gin.H{"errors": errorMessages, "file": zipFileName})
}
