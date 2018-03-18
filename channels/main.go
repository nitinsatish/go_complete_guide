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
	c := make(chan string)

	for _, site := range sites {
		go pingSites(site, c)
	}
	for i := 0; i < len(sites); i++ {
		fmt.Println(<-c)
	}

}

func pingSites(site string, c chan string) {
	_, err := http.Get(site)
	if err != nil {
		fmt.Println(err)
		c <- site + " is down"
		return
	}
	c <- site + " is UP!"
}
