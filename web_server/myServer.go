package main 

import (
	"fmt"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Whoa go is cool")
}

func root_handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hey there!</h1>")
	fmt.Fprintf(w, "<p>Go is so fast</p>")
	fmt.Fprintf(w, "<p>... and simple</p>")
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/go", root_handler)
	http.ListenAndServe(":8000", nil)
}

