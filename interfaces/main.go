package main

type englishBot struct{}
type spanishBot struct{}

func main() {

}

func (eb englishBot) getGreeting() string {
	// VERY custom login for generating an enlish greeting
	return "Hi there"
}

func (sb spanishBot) getGreeting() string {
	return "Hola amigo"
}
