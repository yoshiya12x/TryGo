package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ENDPOINT = "https://api.cognitive.microsoft.com/bing/v7.0/images/search"
	API_KEY  = "hoge"
)

type BingJson struct {
	Type         string `json:"_type"`
	QueryContext struct {
		OriginalQuery string `json:"originalQuery"`
	} `json:"queryContext"`
	Value []struct {
		ContentUrl string `json:"contentUrl"`
	} `json:"value"`
}

func main() {
	flag.Parse()
	searchWord1 := flag.Arg(0)
	//	searchWord2 := flag.Arg(1)
	count := flag.Arg(1)
	execApi(searchWord1, count)
}

func execApi(searchWord string, count string) {
	// Create new http request
	req, err := http.NewRequest("GET", ENDPOINT, nil)
	errorHandling(err)

	// Add get parameters
	params := req.URL.Query()
	params.Add("q", searchWord)
	params.Add("count", count)
	req.URL.RawQuery = params.Encode()

	// Add request header
	req.Header.Add("Ocp-Apim-Subscription-Key", API_KEY)

	// Exec request with new http client
	client := new(http.Client)
	resp, err := client.Do(req)
	errorHandling(err)

	// Close resp
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	errorHandling(err)

	// Parse json
	bingJson := new(BingJson)
	err = json.Unmarshal(body, &bingJson)
	errorHandling(err)

	for i, v := range bingJson.Value {
		fmt.Printf("%d, %s", i, v)
	}
}

func errorHandling(err error) {
	if err != nil {
		panic(err)
	}
}
