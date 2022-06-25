package main

import (
	"fmt"
	"github.com/sunshineplan/imgconv"
	"image"
	"log"
)

type ImgBuildException struct{ Reason string }

func (except *ImgBuildException) Error() string {
	return ""
}

type ImageBuilder struct {
	img         image.Image
	resize      string
	watermark   string
	convertType string
}

func Build(randomFileName, filePath string, resize string, watermark, convertType string) (*ImageBuilder, error) {
	src, err := imgconv.Open(filePath)
	if err != nil {
		log.Println("Error opening file: ", err)
		return nil, &ImgBuildException{"Cannot open file"}
	}

	ib := ImageBuilder{img: src, resize: resize, watermark: watermark, convertType: convertType}
	ib.doResize()
	ib.doWatermark()
	ib.doConvert(randomFileName)

	return &ib, nil
}

func (imgBuilder *ImageBuilder) doResize() {

	imgconv.Resize(imgBuilder.img, imgconv.ResizeOption{Percent: 50})
}

func (imgBuilder *ImageBuilder) doWatermark() {
	imgconv.Watermark(imgBuilder.img, imgconv.WatermarkOption{Mark: markImage, Opacity: 128, Offset: image.Pt(5, 5)})
}

func (imgBuilder *ImageBuilder) doConvert(randomFileName string) error {
	err := imgconv.Save(
		fmt.Sprintf("%v%v.png", ImgResultFolder, randomFileName),
		imgBuilder.img,
		imgconv.FormatOption{Format: imgconv.PNG},
	)
	if err != nil {
		log.Println("Error converting of file: ", err)
		return &ImgBuildException{"Cannot save the converted file"}
	}

	return nil
}
