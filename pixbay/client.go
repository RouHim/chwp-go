package pixbay

import (
	"fmt"
	"github.com/RouHim/chwp/cli"
	"github.com/RouHim/chwp/display"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

var baseUrl = "https://pixabay.com/api/?key=15495421-a5108e860086b11eddaea0efa&per_page=3"

// var baseUrl = "https://pixabay.com/api/?key=15495421-a5108e860086b11eddaea0efa&q=space&min_width=5760&min_height=1080"

// loads a random wallpaper url for given keywords
func GetImageUrl(parameter cli.Configuration, displayInfo display.Information) string {
	requestUrl := baseUrl
	requestUrl += "&q=" + parameter.Keywords[0]
	requestUrl += "&min_width=" + getWidth(displayInfo.TotalResolution)
	requestUrl += "&min_height=" + getHeight(displayInfo.TotalResolution)

	fmt.Println(displayInfo.TotalResolution)
	fmt.Println(requestUrl)

	jsonData := getStringFromUrl(requestUrl)

	imageUrl, _, _, err := jsonparser.Get(jsonData, "hits", "[0]", "largeImageURL")
	//imageUrl, _, _, err := jsonparser.Get(jsonData, "hits", "[0]", "imageURL")

	if err != nil {
		log.Fatal(err)
	}

	return string(imageUrl)
}

func getHeight(resolution string) string {
	return strings.Split(resolution, "x")[1]
}

func getWidth(resolution string) string {
	return strings.Split(resolution, "x")[0]
}

func getStringFromUrl(url string) []byte {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("GET error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Status error: %v", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Read body: %v", err)
	}

	return data
}
