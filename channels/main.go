package main

import (
	"net/http"
	"fmt"
)

func main() {
	links := []string {
		"http://google.com",
		"http://yahoo.com",
		"http://facebook.com",
		"https://golang.io",
		"http://amazon.com",
		"http://creditkarma.com",
	}

	for _, link := range links {
		 checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "is unhealthy!!")
		return
	}

	fmt.Println(link, "is super duper healthy")
}