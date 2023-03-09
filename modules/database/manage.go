package database

import (
	"database/sql"
	"github.com/schollz/progressbar/v3"
	"goScrapper/modules/process"
	"goScrapper/modules/request"
	check "goScrapper/modules/utils"
	"log"
	"strings"
)

func (db *Database) AddCity(name string) {
	sqldb := db.OpenDB()
	defer func(sqldb *sql.DB) {
		err := sqldb.Close()
		check.Err(err)
	}(sqldb)
	insertCity := "INSERT INTO cities(name) VALUES (?)"
	statement, err := sqldb.Prepare(insertCity)
	check.Err(err)
	_, err = statement.Exec(name)
	check.Err(err)
}
func (db *Database) AddStation(station process.Station, city string) {
	sqldb := db.OpenDB()
	defer func(sqldb *sql.DB) {
		err := sqldb.Close()
		check.Err(err)
	}(sqldb)
	insertCity := "INSERT INTO stations(city,brand,imgSrc,address,zipCode,gazole,sp95,e10,sp98,e85,gplc) VALUES (?,?,?,?,?,?,?,?,?,?,?)"
	statement, err := sqldb.Prepare(insertCity)
	check.Err(err)
	_, err = statement.Exec(city, station.Title, station.ImgSrc, station.Address, station.ZipCode, station.Prices[0], station.Prices[1], station.Prices[2], station.Prices[3], station.Prices[4], station.Prices[5])
	check.Err(err)
}

func InsertData(bar *progressbar.ProgressBar, links []string, db *Database) int {
	errCount := 0
	for i := range links {
		content := request.GetHtmlData(links[i])
		city := content.Find("h3").Text()
		city, _, _ = strings.Cut(city, " 01")
		city = strings.ToLower(city)
		city = strings.Title(city)
		db.AddCity(city)
		allStationsInfo := process.ScrapStationInfos(content)
		prices := process.ScrapPricesByCities(content)
		pricesByStation := process.SortPrices(prices)
		cityStations, ok := process.ConvertToObject(allStationsInfo, pricesByStation)
		for j := range cityStations {
			db.AddStation(cityStations[j], city)
		}
		if !ok {
			log.Println("[STEP 4] Exfiltrate data (", i+1, "of", len(links), ") -> Scrapping", city, "[⨯]")
			errCount++
		} else {
			log.Println("[STEP 4] Exfiltrate data (", i+1, "of", len(links), ") -> Scrapping", city, "[✓]")
		}
		err := bar.Add(1)
		check.Err(err)
	}
	return errCount
}
