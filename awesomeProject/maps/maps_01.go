package main

import (
	"fmt"
)

func main() {
	var greetings = make(map[string]string)

	greetings["one"] = "Digambar"
	greetings["two"] = "Ashish"
	greetings["three"] = "Aditya"

	fmt.Println("Greetings map : ", greetings)
	//Delete key from map
	delete(greetings, "three")
	fmt.Println(greetings)

	months := map[int]string{
		0: "January",
		1: "Feb",
		2: "Mar",
		3: "April",
	}
	fmt.Println("Months : ", months)

	delete(months, 2)
	if val, exists := months[2]; exists {
		fmt.Println("Value does exists")
		fmt.Println("Value is ", val)
		fmt.Println("Months are ", months)
	} else {
		fmt.Println("Value doesn't exists")
	}

	if name, ok := months[2]; ok {
		fmt.Println("Value", name, "is", ok)
	} else {
		fmt.Println("Value", name, "is", ok)
	}
}
