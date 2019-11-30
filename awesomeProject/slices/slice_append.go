package main

import (
	"fmt"
)

func main(){
	x := []int{2, 4, 6, 8, 12}
	fmt.Println(x)

	x = append(x, 34, 56, 78)
	fmt.Println(x)

	y := []int{234, 456, 567, 789}
	x = append(x, y...)
	fmt.Println(x)

	//Delete in slice
	x = append(x[:2], x[4:]...)
	fmt.Println(x)

	//Create a slice with make
	a := make([]int, 10, 12)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 123)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 345)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))

	a = append(a, 567)
	fmt.Println(a)
	fmt.Println(len(a))
	fmt.Println(cap(a))
}
