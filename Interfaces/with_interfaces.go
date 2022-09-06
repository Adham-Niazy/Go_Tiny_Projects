package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb1 := englishBot{}
	sb1 := spanishBot{}

	printGreeting(eb1)
	printGreeting(sb1)
}

// Now we condenced the printGreeting Function into one function
func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

// GetGreetings Have different Implementation in both bots
func (englishBot) getGreeting() string {
	// Imagine very custom logic specially for english bot
	return "Hello there"
}
func (spanishBot) getGreeting() string {
	// Imagine very custom logic specially for spanish bot
	return "Hola!"
}