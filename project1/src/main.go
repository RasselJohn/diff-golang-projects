package main

import (
	"img_convert/views"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()
	router.MaxMultipartMemory = 102400 // 100 MB

	router.StaticFile("/", "./static/main.html")
	router.StaticFS("/assets/", http.Dir("./static/assets"))
	router.StaticFS("/public/", http.Dir("./static/public"))
	router.StaticFS("/user_files/", http.Dir("./user_files/result"))

	router.POST("/change-image", views.ChangeImgView)

	// listen on 0.0.0.0:8080
	router.Run()
}
