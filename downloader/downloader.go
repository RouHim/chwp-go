package downloader

import (
	"io/ioutil"
	"log"
	"net/http"
)

func Download(imageUrl string) []byte {
	response, err := http.Get(imageUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	if response.StatusCode != 200 {
		log.Fatal("Received non 200 response code")
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	return bytes
}
