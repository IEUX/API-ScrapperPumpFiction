package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"goScrapper/modules/CLI"
	"goScrapper/modules/database"
	"goScrapper/modules/process"
	"goScrapper/modules/request"
	check "goScrapper/modules/utils"
	"log"
	"os"
	"os/user"
	"time"
)

var db = &database.Database{
	Path:         "database/scrapeer",
	DatabaseName: "scrapperDatabase",
}

func main() {
	//init
	startTime := time.Now()
	startDate := startTime.Format("02.01.2006")
	host, _ := os.Hostname()
	currentUser, _ := user.Current()
	//log setup
	logFile, err := os.OpenFile("logs/local/"+startDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	if !check.Err(err) {
		log.Fatal("[Error] can't open log file")
	}
	log.SetFlags(log.Ltime)
	log.SetOutput(logFile)
	log.Println("[INIT] Scrapper launch by " + currentUser.Username + " on " + host)
	//open database
	db.Init()
	//init done
	//execute request
	url := "https://prix-carburants-info.fr/"
	response, ok, action := request.DoGetRequest(url)
	if !ok {
		log.Println("[STEP 1] GET request [⨯]")
		log.Fatal("[Error]-> Fail to do <", action, "> on : ", url)
	}
	log.Println("[STEP 1] GET request [✓]")
	//open response
	document, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		log.Println("[STEP 2] Open response [⨯]")
		log.Fatal("[Error]-> Fail to <loading HTTP response body> ", err)
	}
	log.Println("[STEP 2] Open response [✓] ")
	//process response
	links, isEmpty := process.ScrapLinks(document)
	if !isEmpty {
		log.Println("[STEP 3] Process response -> Scrap links [⨯]: No links found")
	} else {
		log.Println("[STEP 3] Process response -> Scrap links [✓]:", len(links), "links founded")
	}
	links, ok = process.SortCitiesLinks(links)
	if !ok {
		log.Println("[STEP 3] Process response -> Sort links [⨯]")
	}
	log.Println("[STEP 3] Process response -> Sort links [✓]:", len(links), "cities links kept")
	errCount := 0
	bar := CLI.NewProgressBar(len(links))
	errCount = database.InsertData(bar, links, db)
	//reset log output
	endTime := time.Now()
	totalTime := endTime.Sub(startTime)
	log.Println("[STEP 5] Scrapping done in "+totalTime.String()+" [✓] ( Number of error:", errCount, ")\n<-!->")
	log.SetOutput(os.Stdout)
	//close instances
	err = logFile.Close()
	check.Err(err)
	err = response.Body.Close()
	check.Err(err)
	fmt.Println("")
}
