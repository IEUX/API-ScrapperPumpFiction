package process

import (
	check "goScrapper/modules/utils"
	"net/url"
	"strings"
)

func SortCitiesLinks(links []string) ([]string, bool) {
	checksum := true
	var sortedLinks []string
	for i := range links {
		parsedUrl, err := url.Parse(links[i])
		checksum = check.Err(err)
		splitPath := strings.Split(parsedUrl.Path, "/")
		if len(splitPath) > 1 {
			if splitPath[1] == "villes" {
				sortedLinks = append(sortedLinks, links[i])
			}
		}
	}
	return sortedLinks, checksum
}

func SortPrices(prices []string) [][]string {
	var eachStationPrices [][]string
	for i := 0; i < len(prices); i = i + 6 {
		var stationPrices []string
		for j := 0; j < 6; j++ {
			stationPrices = append(stationPrices, prices[j+i])
		}
		eachStationPrices = append(eachStationPrices, stationPrices)
	}
	return eachStationPrices
}

func ConvertToObject(stations [][]string, prices [][]string) ([]Station, bool) {
	var stationsInfos []Station
	if len(stations) != len(prices) {
		return nil, false
	}
	for i := range stations {
		var stationInfos Station
		stationInfos.Title = stations[i][0]
		stationInfos.ImgSrc = stations[i][1]
		stationInfos.Address = stations[i][2]
		stationInfos.ZipCode = stations[i][3]
		stationInfos.Prices = prices[i]
		stationsInfos = append(stationsInfos, stationInfos)
	}
	return stationsInfos, true
}
