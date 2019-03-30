package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	ENDPOINT = "https://api.cognitive.microsoft.com/bing/v7.0/images/search"
	API_KEY  = "hoge"
)

func main() {
	fmt.Println(ENDPOINT)
	flag.Parse()
	searchWord := flag.Arg(0)
	execApi(searchWord)
	fmt.Println(searchWord)
}

func execApi(searchWord string) {
	// Create new http request
	req, err := http.NewRequest("GET", ENDPOINT, nil)
	errorHandling(err)

	// Add get parameters
	params :=req.URL.Query()
}

func errorHandling(err error) {
	if err != nil{
		panic(err)
	}
}
