package main

import (
	"encoding/xml"
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

/* =======================================  Structs  ===================================== */

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

type NewsAggregatorPage struct {
	Title string
	News  map[string]NewsMap
}

/* ============================================  Methods  ====================================== */

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa go is cool")
}

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there!</h1>")
	fmt.Fprintf(w, "<p>Go is so fast</p>")
	fmt.Fprintf(w, "<p>... and simple</p>")
}

func MyNewsRoutine(channel chan News, LocationURL string) {

	defer wg.Done()
	var story News
	response, _ := http.Get(LocationURL)
	bytes, _ := ioutil.ReadAll(response.Body)
	xml.Unmarshal(bytes, &story)
	response.Body.Close()

	/* Add story to channel */
	channel <- story
}

func news_aggregator_handler(w http.ResponseWriter, r *http.Request) {

	/* ================  Data code  ====================*/
	WaPoUrl := "https://www.washingtonpost.com/news-sitemap-index.xml"

	response, _ := http.Get(WaPoUrl)
	bytes, _ := ioutil.ReadAll(response.Body)
	response.Body.Close()

	var s SiteMapIndex
	xml.Unmarshal(bytes, &s)

	/* Initialize map */
	News_Map := make(map[string]NewsMap)

	/* Initialize channel for news with buffer size = 30 */
	NewsChannel := make(chan News, 30)

	for _, Location := range s.Locations {

		wg.Add(1)
		go MyNewsRoutine(NewsChannel, Location)

	}

	/* Close channel */
	wg.Wait()
	close(NewsChannel)

	/* Populate News_Map from NewsChannel */
	for story := range NewsChannel {
		for idx, _ := range story.Titles {
			News_Map[story.Titles[idx]] = NewsMap{story.Keywords[idx], story.Locations[idx], story.PubDate[idx]}
		}
	}

	p := NewsAggregatorPage{Title: "Amit's awesome sauce news page", News: News_Map}
	t, _ := template.ParseFiles("myNewsApp_template.html")
	t.Execute(w, p)
}

func main() {

	/* =================  Server code  ================ */
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/go", root_handler)
	http.HandleFunc("/news", news_aggregator_handler)
	fmt.Println("Server started at port: 8000")
	http.ListenAndServe(":8000", nil)

}
