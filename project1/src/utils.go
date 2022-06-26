package main

import (
	"./consts"
	"encoding/hex"
	"fmt"
	"math/rand"
	"net/http"
	"strings"
)

func generateFilePath(format string) (randomFileName string, randomFilePath string) {
	randByteData := make([]byte, 16)
	rand.Read(randByteData)
	fileName := hex.EncodeToString(randByteData)
	filePath := fmt.Sprintf("%v%v.%v", consts.ImgLoadFolder, fileName, format)
	return fileName, filePath
}

func getCurrImgFormat(fileName string) string {
	filePathParts := strings.Split(fileName, ".")
	return filePathParts[len(filePathParts)-1]
}

func errorResponse(reason string) (statusCode int, message string) {
	return http.StatusBadRequest, fmt.Sprintf("Error happend. Reason: '%v'. Please, repeat later.", reason)
}
