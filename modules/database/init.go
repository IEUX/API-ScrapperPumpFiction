package database

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"goScrapper/modules/config"
	check "goScrapper/modules/utils"
	"log"
	"os"
)

type Database struct {
	DatabaseName string
	Path         string
}

func (db *Database) Init() {
	//look for a db file
	var dbFile *os.File
	exist, err := check.FileExists("./" + db.Path + "/" + db.DatabaseName + ".db")
	check.Err(err)
	if exist {
		err := os.Remove("./" + db.Path + "/" + db.DatabaseName + ".db")
		check.Err(err)
	}
	dbFile, err = os.Create(db.Path + "/" + db.DatabaseName + ".db")
	check.Err(err)
	log.Println("[DATABASE]Database created")

	//config database
	statements := config.SQLConfig()
	if len(statements) < 1 {
		log.Fatal("[ERR] SQL config file not valid")
	}
	for i := range statements {
		var content string
		for j := range statements[i].Content {
			content += statements[i].Content[j]
		}
		db.CreateTable(statements[i].Name, content)
	}
	err = dbFile.Close()
	check.Err(err)
}
func (db *Database) OpenDB() *sql.DB {
	sqliteDb, err := sql.Open("sqlite3", "./"+db.Path+"/"+db.DatabaseName+".db")
	check.Err(err)
	return sqliteDb
}

func (db *Database) CreateTable(name string, content string) {
	sqldb := db.OpenDB()
	defer func(sqldb *sql.DB) {
		err := sqldb.Close()
		check.Err(err)
	}(sqldb)
	table := "CREATE TABLE IF NOT EXISTS " + name + " (" + content + ");"
	preparation, err := sqldb.Prepare(table)
	check.Err(err)
	_, err = preparation.Exec()
	if err != nil {
		check.Err(err)
	}
}
