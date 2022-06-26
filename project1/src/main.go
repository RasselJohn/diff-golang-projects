package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 20480 // 20 MiB
	router.POST("/change-image", changeImgView)

	// listen on 0.0.0.0:8080
	router.Run()
}
