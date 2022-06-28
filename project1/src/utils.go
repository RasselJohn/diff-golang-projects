package main

import (
	"encoding/hex"
	"fmt"
	"math/rand"
	"strings"
)

func generateFilePath(baseDir, format string) (randomFileName string, randomFilePath string) {
	randByteData := make([]byte, 16)
	rand.Read(randByteData)
	fileName := hex.EncodeToString(randByteData)
	filePath := fmt.Sprintf("%v%v.%v", baseDir, fileName, format)
	return fileName, filePath
}

func getCurrImgFormat(fileName string) string {
	filePathParts := strings.Split(fileName, ".")
	return filePathParts[len(filePathParts)-1]
}
