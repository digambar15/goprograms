package main

import "fmt"

func main(){
	for i := 0; i <= 10; i++{
		fmt.Println(i)
	}

	x :=0
	for {
		x++
		if x > 10 {
			break
		}
		if x % 2 != 0 {
			continue
		}
		fmt.Println(x)
	}

	for i := 65; i <=122; i++{
		fmt.Printf("%v\t%#U\n", i, i)
	}
	}
