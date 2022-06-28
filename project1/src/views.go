package main

import (
	"./consts"
	"./models"
	"archive/zip"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
)


func changeImgView(request *gin.Context) {
	var errorMessages []string

	form, err := request.MultipartForm()
	if err != nil {
		log.Println("Error loading form: ", err)
		errorMessages = append(errorMessages, "Cannot upload file")
		request.JSON(http.StatusBadRequest, gin.H{"errors": errorMessages, "file": ""})
	}

	fileNamesChannel := make(chan string)
	errorChannel := make(chan string)

	files := form.File["img_files[]"]
	for _, file := range files {
		go fileHandler(request, file, fileNamesChannel, errorChannel)
	}

	var fileNames []string
	for i := 0; i < len(files); i++ {
		fileNames = append(fileNames, <-fileNamesChannel)
		errorMessages = append(errorMessages, <-errorChannel)
	}

	zipFileName, zipFilePath := generateFilePath(consts.ImgResultFolder, "zip")
	if err := compressToZip(zipFilePath, fileNames); err != nil {
		request.JSON(http.StatusBadRequest, gin.H{"errors": "Cannot compress to zip", "file": ""})
	}

	request.JSON(http.StatusOK, gin.H{"errors": errorMessages, "file": zipFileName})
}

func fileHandler(request *gin.Context, file *multipart.FileHeader, fileNamesChannel, errorChannel chan string) {
	log.Println("File loaded: ", file.Filename)
	currFormat := getCurrImgFormat(file.Filename)
	fileName, filePath := generateFilePath(consts.ImgLoadFolder, currFormat)

	if err := request.SaveUploadedFile(file, filePath); err != nil {
		log.Println("Error saving of file: ", err)

		fileNamesChannel <- ""
		errorChannel <- fmt.Sprintf("Cannot save file %v", file.Filename)
	}

	size, errorMessage := models.GetNewSize(request.PostForm("width"), request.PostForm("height"))
	if errorMessage != "" {
		fileNamesChannel <- ""
		errorChannel <- errorMessage
	}

	format := models.Format{Old: currFormat, New: request.PostForm("format")}
	newFilePath, errorMessage := models.Build(fileName, filePath, size, format)
	if errorMessage != "" {
		fileNamesChannel <- ""
		errorChannel <- fmt.Sprintf("Error handling of file '%v': %v", file.Filename, errorMessage)
	}

	fileNamesChannel <- newFilePath
	errorChannel <- ""
}

func compressToZip(zipFilePath string, fileNames []string) error {
	zipFile, err := os.Create(zipFilePath)
	if err != nil {
		return err
	}
	defer zipFile.Close()

	zipWriter := zip.NewWriter(zipFile)
	defer zipWriter.Close()

	for _, fileName := range fileNames {
		if fileName == "" {
			continue
		}

		file, err := os.Open(fileName)
		if err != nil {
			return err
		}
		defer file.Close()

		writer, err := zipWriter.Create(fileName)
		if err != nil {
			return err
		}

		if _, err := io.Copy(writer, file); err != nil {
			return err
		}
	}

	return nil
}
