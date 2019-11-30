package main

import "fmt"

func max(num ...int) (int){
	var large int
	for _, v := range num{
		if v > large {
			large = v
		}
	}
	return large
}

func main(){
	greatest := max(23, 21, 24, 43, 65, 98)
	fmt.Println(greatest)
}
