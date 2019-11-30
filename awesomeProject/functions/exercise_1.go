package main

import (
	"fmt"
)

func cal(n int) (int, bool){
	div := n/2
	var mod bool
	if n%2 == 0 {
		mod = true
	}

	if div != 0 && mod {
		return div, mod
	}
	return 0, false
}

func main(){
	div, st := cal(5)
	fmt.Println(div, st)
}