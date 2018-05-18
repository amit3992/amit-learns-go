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
	//fmt.Println("Here are the news links from Washington Post: \n")
	count := 0
	
	for _, Location := range s.Locations {
		//fmt.Printf("%d] %s\n", count, Location)
		
		response, _ := http.Get(Location)
		bytes, _ := ioutil.ReadAll(response.Body)
		response.Body.Close()
		xml.Unmarshal(bytes, &story)
		count++
	}
	
	fmt.Println("Here are the latest stories in Politics from Washington Post: \n")
	storyCount := 0
	for _, Title := range story.Titles {
		fmt.Printf("%d] %s\n", storyCount, Title)
		storyCount++
	}
	
	
	
}