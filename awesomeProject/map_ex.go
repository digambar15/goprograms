package main

import "fmt"

func main() {
	m := map[string]int{}

	fmt.Println("map = ", m)
	m["K1"] = 1
	m["K2"] = 2
	fmt.Println(m)

	//Delete key from map
	delete(m, "K2")
	fmt.Println(m)

	if v, ok := m["K1"]; ok {
		fmt.Println("Value exist ", ok, v)
	}else{
		fmt.Println("Value doesn't exist")
	}


	n := map[string]int{"Welcome": 2019}
	fmt.Println("Welcome to ", n["Welcome"])

	//Declare map using make function
	mp := make(map[string]string)

	mp["Pune"] = "Cultural City"
	mp["Mumbai"] = "Finance Capital of India"
	for key, val := range mp{
		fmt.Println(key, " = ", val)
	}
	//var nested_map = make(map[int]map[string]string)
}

