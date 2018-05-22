package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type NewsAggregatorPage struct {
	Title string
	News  string
}

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa go is cool")
}

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there!</h1>")
	fmt.Fprintf(w, "<p>Go is so fast</p>")
	fmt.Fprintf(w, "<p>... and simple</p>")
}

func news_aggregator_handler(w http.ResponseWriter, r *http.Request) {
	p := NewsAggregatorPage{Title: "My awesome sauce news page", News: "some news"}
	t, _ := template.ParseFiles("sexy_template.html")
	t.Execute(w, p)
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/go", root_handler)
	http.HandleFunc("/news", news_aggregator_handler)
	http.ListenAndServe(":8000", nil)
}
