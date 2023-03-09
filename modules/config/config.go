package config

import (
	"encoding/json"
	check "goScrapper/modules/utils"
	"io/ioutil"
	"log"
	"os"
)

type Statement struct {
	Name    string   `json:"name"`
	Content []string `json:"content"`
}

func SQLConfig() []Statement {
	var statements []Statement
	configFile, err := os.Open("config/sqlConfig.json")
	check.Err(err)
	log.Println("[JSON] Statement file successfully open")
	defer func(configFile *os.File) {
		err := configFile.Close()
		check.Err(err)
	}(configFile)
	fileContent, _ := ioutil.ReadAll(configFile)
	err = json.Unmarshal(fileContent, &statements)
	check.Err(err)
	return statements
}
