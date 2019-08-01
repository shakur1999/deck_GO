package main

import "fmt"

func main() {
	var colors := make(map[string] string) // this is to declare a colors map as string to string key value pair  
	colors := map[int]string{
		"red":   "#12345",
		"green": "#54321",
	}
	colors["white"] = "#45678" //this is to add a new string map value 
	colors[10] = "#09876" // this is to add a new int map value
	delete(colors, 10) // this is to delete a value in colors map

	fmt.Println(colors)
}
