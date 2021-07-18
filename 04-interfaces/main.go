package main

import "fmt"

type englishBot struct{}

type spanishBot struct{}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating english greeting
	return "Hi There!"
}

func printGreeting(eb englishBot) {
	fmt.Println(eb.getGreeting())
}

func (spanishBot) getGreeting() string {
	// VERY custom logic for generating spanish greeting
	return "Hola!"
}

func printGreeting(sb spanishBot) {
	fmt.Println(sb.getGreeting())
}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}
