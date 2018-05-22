package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

/* Structs */
type SiteMapIndex struct {
	Locations []string `xml:"sitemap>loc"`
}

type News struct {
	Titles    []string `xml:"url>news>title"`
	Keywords  []string `xml:"url>news>keywords"`
	PubDate   []string `xml:"url>news>publication_date"`
	Locations []string `xml:"url>loc"`
}

type NewsMap struct {
	Keyword         string
	Location        string
	PublicationDate string
}

/* Methods */
func (s SiteMapIndex) showLocations() {

	fmt.Println("Here are the locations: \n")
	count := 0
	for _, Location := range s.Locations {
		fmt.Printf("%d] %s\n", count, Location)
		count++
	}
}

func main() {

	/* Get info from the internet */
	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")

	bytes, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	var story News
	News_Map := make(map[string]NewsMap)

	for _, Location := range s.Locations {
		//fmt.Printf("%d] %s\n", count, Location)

		response, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		xml.Unmarshal(bytes, &story)

		for idx, _ := range story.Titles {
			News_Map[story.Titles[idx]] = NewsMap{story.Keywords[idx], story.Locations[idx], story.PubDate[idx]}
		}
	}

	for index, data := range News_Map {
		fmt.Printf("Index \t Location \t Publication Data\n")
		fmt.Printf("%s\t%s\t%s\t%s\n", index, data.Location, data.PublicationDate)
	}

}
