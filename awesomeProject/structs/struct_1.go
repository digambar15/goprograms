package main

import "fmt"

type person struct {
	first string
	last string
	age int
}

func main() {
	p1 := person{"Digambar", "Patil", 33}
	p2 := person{"Varad", "Patil", 4}
	fmt.Println("First candidate = ", p1)
	fmt.Println("First candidate = ", p2)

}

