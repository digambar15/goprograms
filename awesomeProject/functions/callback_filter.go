package main

import "fmt"

func filter(numbers []int, callbacks func(int) bool) []int{
	xs := []int{}
	for _, n := range numbers{
		if callbacks(n){
			xs = append(xs, n)
		}
	}
	return xs
}

func main(){
	xs := filter([]int{10, 13, 14, 30, 54, 45, 33, 50}, func(n int) bool{
		return n % 2 == 0
	})
	fmt.Println(xs)
}
