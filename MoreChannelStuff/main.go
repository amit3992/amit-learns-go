package main

import (
	"fmt"
	"net/http"
)

func checkLink(link string) {

	//fmt.Printf("Checking link -> %s\n", link)
	_, err := http.Get(link)

	if err != nil {
		fmt.Printf("%s is down :(. Error -> %s\n", link, err)

	} else {
		fmt.Printf("%s is up!\n", link)
	}
}

func main() {

	links := []string{
		"https://google.com",
		"https://facebook.com",
		"https://stackoverflow.com",
		"https://golang.org",
		"https://amazon.com",
	}

	for _, link := range links {
		checkLink(link)
	}
}
