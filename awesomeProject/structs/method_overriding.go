package main

import "fmt"

type PersonData struct {
	first string
	last  string
	age   int
}

type DoubleData struct {
	PersonData
	first string
	last  string
}

func (p PersonData) Greeting() {
	fmt.Println("I am a regular class")
}

func (dz DoubleData) Greeting() {
	fmt.Println("This is extended class")
}

func main() {
	p1 := PersonData{"Noopur", "Landge", 27}
	p2 := DoubleData{
		PersonData: PersonData{
			first: "Girish",
			last:  "Landge",
			age:   33,
		},
		first: "Amey",
		last:  "Raynade",
	}
	p1.Greeting()
	p2.Greeting()
	p2.PersonData.Greeting()
}
