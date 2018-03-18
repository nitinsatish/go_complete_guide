package main

import (
	"fmt"
	"net/http"
)

func main() {
	sites := []string{
		"http://google.com",
		"http://golang.org",
		"http://stackoverflow.com",
		"http://reddit.com",
		"http://twitter.com",
		"http://amazon.com",
	}
	for _, site := range sites {
		checkSites(site)
	}
}

func checkSites(site string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(site, "is down!")
		return
	}
	fmt.Println(site, "is UP!")
}
