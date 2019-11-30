package main

import "fmt"

func main(){
	x := []int{2,3,4,5,67,8}
	fmt.Println(len(x))
	fmt.Println(x)

	for i, v := range x{
		fmt.Println(i, v)
	}
}
