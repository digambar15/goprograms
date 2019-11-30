package main

import "fmt"

func main() {
	//Printout number which is found between 10 to 100 when it is divided by 4.
	for i := 10; i <= 100; i++ {
		if i % 4 == 0 {
			fmt.Println(i)
		}
	}

	x := "Hello"
	if x == "Hellow" {
		fmt.Println(x)
	} else if x == "Hello" {
		fmt.Println("Hello")
	} else {
		fmt.Println("Done!")
	}

	vv := "FavSport"
	switch vv {
	case "AIM":
		fmt.Println("Not correct")
	case "FavSport":
		fmt.Println("You found it")
	default:
		fmt.Println("Done!")
	}
}