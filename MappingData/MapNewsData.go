package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

/* Structs */
type SiteMapIndex struct {
	Locations [] string `xml:"sitemap>loc"`
}

type News struct {
	Titles [] string `xml:"url>news>title"`
	Keywords [] string `xml:"url>news>keywords"`
	PubDate [] string `xml:"url>news>publication_date"`
	Locations [] string `xml:"url>loc"`
}

type NewsMap struct {
  Keyword string
  Location string
  PubDate string
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

  //fmt.Println("Here are the news links from Washington Post: \n")
	count := 0
	for _, Location := range s.Locations {
		//fmt.Printf("%d] %s\n", count, Location)

		response, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		xml.Unmarshal(bytes, &story)
		count++

    for index, _ := range story.Titles {
      News_Map[story.Titles[index]] = NewsMap{story.Keywords[index], story.Locations[index], story.PubDate[index]}
    }
	}

  for idx, data := range News_Map {
    fmt.Println("\n\n\n", idx)
    fmt.Println("\n", data.Keyword)
    fmt.Println("\n", data.PubDate)
    fmt.Println("\n", data.Location)
  }


}
