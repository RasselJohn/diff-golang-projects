package models

import (
	"../consts"
	"fmt"
	"github.com/disintegration/imaging"
	"image"
	"log"
)

type ImgBuilder struct {
	img    image.Image
	size   *Size
	format Format
}

func Build(fileName, filePath string, size *Size, format Format) (string, string) {
	img, err := imaging.Open(filePath)
	if err != nil {
		log.Println("Error opening file: ", err)
		return "", "Cannot open file"
	}

	imageBuilder := ImgBuilder{img: img, size: size, format: format}
	imageBuilder.resize()

	newPath, errorMessage := imageBuilder.convert(fileName)
	return newPath, errorMessage
}

func (imgBuilder *ImgBuilder) resize() {
	if imgBuilder.size == nil {
		return
	}
	imgBuilder.img = imaging.Resize(imgBuilder.img, imgBuilder.size.width, imgBuilder.size.height, imaging.Lanczos)
}

func (imgBuilder *ImgBuilder) convert(fileName string) (string, string) {
	format := imgBuilder.format.Old
	if imgBuilder.format.New != "" && imgBuilder.format.New != imgBuilder.format.Old {
		format = imgBuilder.format.New
	}

	newPath := fmt.Sprintf("%v%v.%v", consts.ImgResultFolder, fileName, format)
	err := imaging.Save(imgBuilder.img, newPath)
	if err != nil {
		log.Println("Error saving of the handled file: ", err)
		return "", "Cannot save the handled file"
	}

	return newPath, ""
}
