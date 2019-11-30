package main

import (
	"fmt"
)

func main() {
	greeting := []string{
		"Digambar",
		"Pratibha",
		"Varad",
		"Patil",
		"Swara",
		"Manu",
	}

	for i, v := range greeting {
		fmt.Println(i, v)
	}

	for j := 0; j < len(greeting); j++ {
		fmt.Println(greeting[j])
	}

	fmt.Println("Slices are [2:4] ", greeting[2:4])
	fmt.Println("[5:] ", greeting[5:])
	fmt.Println("[:5] ", greeting[:5])

	//Appending a slice to slice
	newSlice := []string{
		"Tasgaon",
		"Sangli",
		"Maharashtra",
		"India",
	}
	greeting = append(greeting, newSlice...)
	fmt.Println("Appended slice is ", greeting)

	records()
	transactions()
}
