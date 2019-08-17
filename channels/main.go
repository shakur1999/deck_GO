package main

import "fmt"

func main() {
	links := [] string {
		"https://google.com",
		"https://yahoo.com",
		"https://facebook.com",
		"https://golang.org",
		"https://amazon.com",
		"https://creditkarma.com"
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

	fmt.Println(link, is super duper healthy)
}