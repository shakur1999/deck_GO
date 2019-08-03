package main

import "fmt"

type bot interface {
	getGreeting() string
}

// option1, we used interfaces to define a method set aka getGreeting() string

type englishBot struct{}
type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom login for generating an enlish greeting
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola amigo!"
}
