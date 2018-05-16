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

func testLoop() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func testWhileLikeLoop() {
	x := 5
	
	for {
		fmt.Println("Do stuff ", x)
		x += 3
		if x > 25 {
			break
		}
	}
}

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
	//string_body := string(bytes)
	//fmt.Println(string_body)
	response.Body.Close()
	
	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)
	
	//fmt.Println(s.Locations)
	s.showLocations()
	//testLoop()
	//testWhileLikeLoop()
	
}