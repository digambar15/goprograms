package main

import "fmt"

func main(){
	var arr [58]string
	fmt.Println(len(arr))

	for i:=65; i<=122; i++{
		arr[i-65] = string(i)
	}

	var candies [20]int
	for i:=1; i<20; i++{
		candies[i] = i
	}

	var elephant [10]int
	for i:=1;i<10;i++{
		elephant[i] = i
	}
	for j:=1;j<10;j++{
		candies[j] = candies[j] - elephant[j]
	}
	fmt.Println(candies)
	fmt.Println(len(candies))
}
