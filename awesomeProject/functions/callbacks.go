package main

import "fmt"

func visit(num []int, callback func(int)){
	for _, n := range num{
		callback(n)
	}
}

func main()  {
	visit([]int{1,2,3,4}, func(n int) {
		n++
		fmt.Println(n)
	})
}