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

	c := make(chan string) //this is how we create a brad new channel in Go

	for _, link := range links {
		 go checkLink(link, c)
		//  fmt.Println(<- c) this seems to do the samething as the FOR loop below
	}

	for {
		go checkLink(<-c, c)
	}
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link) 
		if err != nil {
			fmt.Println(link, "server is DOWN!")
			c <- link
			return
		}
		
		fmt.Println(link, "Server Operational!")
		c <- link
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