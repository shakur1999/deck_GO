package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#12345",
		"green": "#54321",
		"white": "#fffff",
	}

	printMap(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, hex)
	}
}
