package main

import (
	"fmt"
	"math"
)

func main(){
	var n float64 = 19
	for i:=2; i < int(math.Sqrt(n)); i++{
		if (int(n) % i) == 0 {
			fmt.Println("%d is not a prime number", n)
		} else {
			fmt.Println("%d is a prime number", n)
		}
	}
}
