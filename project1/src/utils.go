package main

import (
	"./consts"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
)

func generateFilePath() (randomFileName string, randomFilePath string) {
	randByteData := make([]byte, 16)
	rand.Read(randByteData)
	fileName := hex.EncodeToString(randByteData)
	filePath := fmt.Sprintf("%v%v.jpg", consts.ImgLoadFolder, fileName)
	return fileName, filePath
}

func getFormat(newFormat string) (string, string) {
	format := ""
	if newFormat != "" {
		if newFormat == consts.JPG || newFormat == consts.PNG {
			format = newFormat
		} else {
			return "", "Unknown format file or converting format"
		}
	}

	return format, ""
}

func errorResponse(reason string) (statusCode int, message string) {
	return http.StatusBadRequest, fmt.Sprintf("Error happend. Reason: '%v'. Please, repeat later.", reason)
}
