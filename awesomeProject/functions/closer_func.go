package main

import "fmt"

func main(){
	var x int

	increament := func() int {
		x++
		return x
	}

	fmt.Println(increament())
	fmt.Println(increament())
}