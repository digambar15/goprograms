package main

import (
	"fmt"
)

func main() {

	var val [58]string
	fmt.Println("Size of array is ", len(val))
	fmt.Println("val array = ", val)

	for i := 65; i <= 122; i++ {
		val[i-65] = string(i)
	}
	fmt.Println(val)
	fmt.Println("Size of array is ", len(val))
}
