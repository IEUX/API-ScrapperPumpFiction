package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	request "goScrapper/API/router"
	"goScrapper/config"
	"goScrapper/modules/database"
	check "goScrapper/modules/utils"
	"log"
	"os"
	"os/user"
	"time"
)

var db = &database.Database{
	Path:         "./database/scrapper",
	DatabaseName: "scrapperDatabase",
}

func main() {
	//init
	startTime := time.Now()
	startDate := startTime.Format("02.01.2006")
	host, _ := os.Hostname()
	currentUser, _ := user.Current()
	//prepare logs
	logFile, err := os.OpenFile("logs/server/"+startDate, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0600)
	
	if !check.Err(err) {
		log.Fatal("[Error] can't open log file")
	}
	gin.DefaultWriter = logFile
	gin.SetMode(gin.ReleaseMode)
	log.SetOutput(logFile)
	log.SetFlags(log.Ltime)
	log.Println("[SERVER] Server started by ", currentUser.Username, " on ", host)
	//Gin config
	router := config.CustomRouter()
	request.GetFunctions(router, db)
	port := "8080"
	fmt.Printf("\033[0;36m%s\033[0m", "[SERVER] http://localhost:"+port)
	err = router.Run(":" + port)
	check.Err(err)
}
