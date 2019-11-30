package main

import (
	"fmt"
)

func main()  {
	s := average(12, 23, 34, 45, 56, 67)
	fmt.Println(s)

	ss:= []float64{12.4, 14.5, 16.5, 17.8}
	aa := ages(ss)
	fmt.Println(aa)
}

func average(sf ...float64) (float64)  {
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	var total float64
	for _, val := range sf {
		total += val
	}
	return total / float64(len(sf))
}

func ages(sf []float64) (float64){
	fmt.Println(sf)
	fmt.Printf("%T\n", sf)
	var total float64
	for _, val := range sf {
		total += val
	}
	return total / float64(len(sf))
}