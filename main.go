package main

import (
	"fmt"
	"os"

	"github.com/RouHim/chwp/cli"
	"github.com/RouHim/chwp/display"
	"github.com/RouHim/chwp/pixbay"
)

func main() {
	args := os.Args[1:]
	configuration := cli.Parse(args)
	displayInfo := display.GetDisplayInfo()

	wallpaper := pixbay.GetImageUrl(configuration, displayInfo)

	fmt.Println(wallpaper)
}
