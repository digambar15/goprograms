package main

import (
	"fmt"
)

func multiple_3_and_5(n int) (int){
	var sum int
	for i:=0; i<n; i++{
		if i%3 == 0 || i%5 == 0 {
			sum = sum + i
		}
	}
	return sum
}

func main()  {
	ss := multiple_3_and_5(1000)
	fmt.Println(ss)
}