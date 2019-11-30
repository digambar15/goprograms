package main

import "fmt"

func Linearsort(l []int, callback func(n int) bool) int{
	var i int
	for _, n := range l{
		if callback(n){
			i = n
		}
	}
	return i
}

func main(){
	x := 3
	list := []int{1,2,3,4,5,6}
	xs := Linearsort(list, func(n int) bool{
		return n == x
	})

	if xs != 0 {
		fmt.Printf("Item %d found in %v", xs, list)
	}else {
		fmt.Println("Item not found in the list")
	}
}