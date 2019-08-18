package main

import(
	"net/http"
	"fmt"
)

func main() {
	links := [] string {
		"https://google.com",
		"https://twitter.com",
		"https://facebook.com",
		"https://verizon.com",
		"https://golang.org",
		"https://creditkarma.com",
		"http://code.corp.creditkarma.com",
	}

	for _, link := range links {
		 go checkLink(link)
	}
}

func checkLink(link string) {
	_, err := http.Get(link) 
		if err != nil {
			fmt.Println(link, "server is DOWN!")
			return
		}
		
		fmt.Println(link, "Server Operational!")
}

// package main

// import (
// 	"net/http"
// 	"fmt"
// )

// func main() {
// 	links := []string {
// 		"http://google.com",
// 		"http://yahoo.com",
// 		"http://facebook.com",
// 		"https://golang.io",
// 		"http://amazon.com",
// 		"http://creditkarma.com",
// 	}

// 	for _, link := range links {
// 		 checkLink(link)
// 	}
// }

// func checkLink(link string) {
// 	_, err := http.Get(link)
// 	if err != nil {
// 		fmt.Println(link, "is unhealthy!!")
// 		return
// 	}

// 	fmt.Println(link, "is super duper healthy")
// }