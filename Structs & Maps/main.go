package main

import "fmt"

// Defining Structs
type contactInfo struct {
	email   string
	zipCode int
}
type person struct {
	firstName string
	lastName  string
	contact   contactInfo
}

func main() {
	// Using Structs
	p1 := person{
		firstName: "Adham", 
		lastName: "Niazy",
		contact: contactInfo{
			email: "adhamniazy@gmail.com",
			zipCode: 00000,
		},
	}
	p1.updatedName("John")
	p1.print()


	// Using Maps
	colors := map[string]string {
		"red": "#F00",
		"green": "#0F0",
		"blue": "#00F",
		"black": "#000",
	}

	colors["white"] = "#FFF"
	delete(colors, "black")

	printMap(colors)
}

func (p person) print() {
	fmt.Printf("%+v", p)
}

func (p *person) updatedName(newName string) {
	p.firstName = newName
}

func printMap(m map[string]string) {
	for key, value := range m {
		fmt.Println(key + " => " + value)
	}
}
