package utils

import (
	"img_convert/consts"
	"img_convert/models"
	"archive/zip"
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"math/rand"
	"mime/multipart"
	"os"
	"strings"
)

func FileHandler(request *gin.Context, file *multipart.FileHeader, fileNamesChannel, errorChannel chan string) {
	log.Println("File loaded: ", file.Filename)

	currFormat := getCurrImgFormat(file.Filename)
	if currFormat != "jpg" && currFormat != "png" {
		fileNamesChannel <- ""
		errorChannel <- fmt.Sprintf("File %v has unknown format.", file.Filename)
		return
	}

	fileName, filePath := GenerateFilePath(consts.ImgLoadFolder, currFormat)
	if err := request.SaveUploadedFile(file, filePath); err != nil {
		log.Println("Error saving of file: ", err)

		fileNamesChannel <- ""
		errorChannel <- fmt.Sprintf("Cannot save file %v.", file.Filename)
		return
	}

	size, errorMessage := models.GetNewSize(request.PostForm("width"), request.PostForm("height"))
	if errorMessage != "" {
		fileNamesChannel <- ""
		errorChannel <- errorMessage
		return
	}

	format := models.Format{Old: currFormat, New: request.PostForm("format")}
	newFilePath, errorMessage := models.Build(fileName, filePath, size, format)
	if errorMessage != "" {
		fileNamesChannel <- ""
		errorChannel <- fmt.Sprintf("Error handling of file '%v': %v.", file.Filename, errorMessage)
		return
	}

	fileNamesChannel <- newFilePath
	errorChannel <- ""
}

func getCurrImgFormat(fileName string) string {
	filePathParts := strings.Split(fileName, ".")
	return filePathParts[len(filePathParts)-1]
}

func GenerateFilePath(baseDir, format string) (randomFileName string, randomFilePath string) {
	randByteData := make([]byte, 16)
	rand.Read(randByteData)
	fileName := hex.EncodeToString(randByteData)
	filePath := fmt.Sprintf("%v%v.%v", baseDir, fileName, format)
	return fileName, filePath
}

func CompressToZip(zipFilePath string, fileNames []string) error {
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
