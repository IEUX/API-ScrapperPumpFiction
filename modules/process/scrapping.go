package process

import (
	"github.com/PuerkitoBio/goquery"
	"strings"
)

type Station struct {
	Title   string
	ImgSrc  string
	Address string
	ZipCode string
	Prices  []string
	Gazole  string
}

func ScrapLinks(document *goquery.Document) ([]string, bool) {
	var links []string
	document.Find("a").Each(func(i int, element *goquery.Selection) {
		newTag, exists := element.Attr("href")
		if exists {
			links = append(links, newTag)
		}
	})
	if len(links) == 0 {
		return nil, false
	}
	return links, true
}

func ScrapStationInfos(document *goquery.Document) [][]string {
	var allStationsInfo [][]string
	document.Find(".row").Each(func(i int, selection *goquery.Selection) {
		var stationInfo []string
		brand := selection.Find(".text-decoration-none.text-dark").Text()
		imgSrc, _ := selection.Find("img").Attr("src")
		element := selection.Find(".text-info")
		address, _, _ := strings.Cut(element.Parent().Text(), element.Text())
		addressWithoutBackspace := strings.TrimPrefix(address, "\n")
		zipCode := element.Text()
		stationInfo = append(stationInfo, brand)
		stationInfo = append(stationInfo, imgSrc)
		stationInfo = append(stationInfo, addressWithoutBackspace)
		stationInfo = append(stationInfo, zipCode)
		allStationsInfo = append(allStationsInfo, stationInfo)
	})
	allStationsInfo = allStationsInfo[1:]
	return allStationsInfo
}

func ScrapPricesByCities(document *goquery.Document) []string {
	var allPrices []string
	document.Find(".col-2.text-center.px-0").Each(func(index int, element *goquery.Selection) {
		allPrices = append(allPrices, element.Text())
	})
	return allPrices
}
