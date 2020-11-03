package pixbay

import (
	"fmt"
	"github.com/RouHim/chwp/cli"
	"github.com/RouHim/chwp/display"
	"github.com/buger/jsonparser"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var baseUrl = "https://pixabay.com/api/?key=15495421-a5108e860086b11eddaea0efa&per_page=50"
// var baseUrl = "https://pixabay.com/api/?key=15495421-a5108e860086b11eddaea0efa&q=space&min_width=5760&min_height=1080"

// loads a random wallpaper url for given keywords
func GetImageUrl(parameter cli.Configuration, displayInfo display.Information) string {
	requestUrl := buildRequestUrl(parameter, displayInfo)

	fmt.Println(requestUrl)

	jsonData := getStringFromUrl(requestUrl)

	var images []string
	jsonparser.ArrayEach(jsonData, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		imageUrl, err := jsonparser.GetString(value, "imageURL")

		if err != nil {
			log.Fatal(err)
		}
		images = append(images, imageUrl)
	}, "hits")

	rand.Seed(time.Now().UnixNano())
	randomIndex := rand.Intn(len(images) - 1)
	return images[randomIndex]
}

func buildRequestUrl(parameter cli.Configuration, displayInfo display.Information) string {
	targetWidth := getWidth(displayInfo.MaxSingleResolution)
	if parameter.Span {
		targetWidth = getWidth(displayInfo.TotalResolution)
	}

	targetHeight := getHeight(displayInfo.MaxSingleResolution)
	if parameter.Span {
		targetHeight = getHeight(displayInfo.TotalResolution)
	}

	requestUrl := baseUrl
	requestUrl += "&q=" + parameter.Keywords[0]
	requestUrl += "&min_width=" + targetWidth
	requestUrl += "&min_height=" + targetHeight
	return requestUrl
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
