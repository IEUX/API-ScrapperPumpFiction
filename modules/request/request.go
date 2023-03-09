package request

import (
	"github.com/PuerkitoBio/goquery"
	check "goScrapper/modules/utils"
	"io"
	"log"
	"net/http"
	"time"
)

func NewClient() http.Client {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	return *client
}

func NewRequest(url string) http.Request {
	request, err := http.NewRequest("GET", url, nil)
	check.Err(err)
	request.Header.Set("User-Agent", "Firefox")
	return *request
}

func DoGetRequest(url string) (*http.Response, bool, string) {
	//initials values
	checkSum := true
	client := NewClient()
	request := NewRequest(url)
	//execute request
	response, err := client.Do(&request)
	checkSum = check.Err(err)
	return response, checkSum, "GET request"
}

func GetHtmlData(url string) *goquery.Document {
	DoGetRequest(url)
	response, ok, action := DoGetRequest(url)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		check.Err(err)
	}(response.Body)
	if !ok {
		log.Fatal("[Error]-> Fail to do <", action, "> on : ", url)
	}
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Fatal("[Error]-> Fail to <loading HTTP response body> ", err)
	}
	return document
}
