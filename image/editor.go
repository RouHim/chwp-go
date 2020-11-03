package image

import (
	"bytes"
	"github.com/RouHim/chwp/display"
	"github.com/disintegration/imaging"
	"image"
	"log"
	"strconv"
	"strings"
)

func getImageConfig(imageData *[]byte) image.Config {
	imageConfig, _, err := image.DecodeConfig(bytes.NewReader(*imageData))
	if err != nil {
		log.Fatal(err)
	}
	return imageConfig
}

func ScaleToFitDisplay2(imageData *[]byte, span bool, displayInfo display.Information) []byte {
	displayWidth := getWidth(displayInfo.MaxSingleResolution)
	if span {
		displayWidth = getWidth(displayInfo.TotalResolution)
	}

	displayHeight := getHeight(displayInfo.MaxSingleResolution)
	if span {
		displayHeight = getHeight(displayInfo.TotalResolution)
	}
	displayRatio := float64(displayWidth) / float64(displayHeight)

	imageConfig := getImageConfig(imageData)
	imgWidth := imageConfig.Width
	imgHeight := imageConfig.Height
	imgRatio := float64(imgWidth) / float64(imgHeight)
	srcImage, err := imaging.Decode(bytes.NewReader(*imageData), imaging.AutoOrientation(true))
	if err != nil {
		log.Fatal(err)
	}

	var targetImageWidth int
	var targetImageHeight int
	if imgRatio <= displayRatio {
		targetImageWidth = imgWidth
		targetImageHeight = int(float64(imgWidth) / displayRatio)
	} else {
		targetImageWidth = int(float64(imgHeight) * displayRatio)
		targetImageHeight = imgHeight
	}

	croppedImage := imaging.CropCenter(srcImage, targetImageWidth, targetImageHeight)

	var buf bytes.Buffer
	err = imaging.Encode(&buf, croppedImage, imaging.PNG)
	return buf.Bytes()
}

func getWidth(resolution string) int {
	width, err := strconv.Atoi(strings.Split(resolution, "x")[0])
	if err != nil {
		log.Fatal(err)
	}
	return width
}

func getHeight(resolution string) int {
	height, err := strconv.Atoi(strings.Split(resolution, "x")[1])
	if err != nil {
		log.Fatal(err)
	}
	return height
}
