package database

import (
	"database/sql"
	check "goScrapper/modules/utils"
)

type Station struct {
	id      int
	City    string
	Title   string
	ImgSrc  string
	Address string
	ZipCode string
	Gazole  string
	Sp95    string
	E10     string
	Sp98    string
	E85     string
	Gplc    string
}

func (db *Database) GetStation(query string, value string) []Station {
	sqldb := db.OpenDB()
	var allStations []Station
	defer func(sqldb *sql.DB) {
		err := sqldb.Close()
		check.Err(err)
	}(sqldb)
	rows, err := sqldb.Query("SELECT * FROM stations WHERE " + query + " LIKE '" + value + "'")
	check.Err(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		check.Err(err)
	}(rows)
	for rows.Next() {
		var station Station
		err := rows.Scan(&station.id, &station.City, &station.Title, &station.ImgSrc, &station.Address, &station.ZipCode, &station.Gazole, &station.Sp95, &station.E10, &station.Sp98, &station.E85, &station.Gplc)
		check.Err(err)
		allStations = append(allStations, station)
	}
	return allStations
}
