package main

import (
	"fmt"
	"github.com/RouHim/chwp/downloader"
	"github.com/RouHim/chwp/image"
	"github.com/RouHim/chwp/kde"
	"os"

	"github.com/RouHim/chwp/cli"
	"github.com/RouHim/chwp/display"
	"github.com/RouHim/chwp/pixbay"
)

func main() {
	args := os.Args[1:]
	configuration := cli.Parse(args)
	displayInfo := display.GetDisplayInfo()
	wallpaperUrl := pixbay.GetImageUrl(configuration, displayInfo)
	fmt.Println(wallpaperUrl)
	imageData := downloader.Download(wallpaperUrl)
	imageData = image.ScaleToFitDisplay2(&imageData, configuration.Span, displayInfo)
	kde.ChangeWallpaper(&imageData)
}
