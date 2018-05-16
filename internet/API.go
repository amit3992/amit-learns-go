package main 

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"encoding/xml"
)

type SiteMapIndex struct {
	Locations []Location `xml:"sitemap"`
}

type Location struct {
	Loc string `xml:"loc"`
}

func (l Location) String() string {
	return fmt.Sprintf(l.Loc)
} 

func main() {
	
	/* Get info from the internet */
	response, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
	
	bytes, _ := ioutil.ReadAll(response.Body)
	//string_body := string(bytes)
	//fmt.Println(string_body)
	response.Body.Close()
	
	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)
	
	fmt.Println(s.Locations)
	
}

