package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Datasource Type
type Datasource struct {
	Buyers       string
	Products     string
	Transactions string
}

func GetDataPaths() Datasource {
	path := "repo.json"
	jsonFile, err := os.Open(path)

	if err != nil {
		fmt.Println(err)
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)

	var urls Datasource

	err = json.Unmarshal(byteValue, &urls)

	if err != nil {
		fmt.Println("Error", err)
	}

	return urls
}

// ExtractDataFrom receives a url to fetch
func ExtractDataFrom(url string, date string) ([]byte, error) {
	// Search in repo.json the urls
	urls := GetDataPaths()

	var (
		response *http.Response
		err      error
	)

	if url == "buyers" {
		response, err = http.Get(urls.Buyers + "?date=" + date)
	} else if url == "products" {
		response, err = http.Get(urls.Products + "?date=" + date)
	} else {
		response, err = http.Get(urls.Transactions + "?date=" + date)
	}

	// Managing errors
	if err != nil {
		fmt.Println(err)
	}

	// Close Body on finish scoped function
	defer response.Body.Close()

	// Read the body and manage errors
	return ioutil.ReadAll(response.Body)
}
