package main

import "fmt"

func main(){
	n := HashFunction("Go", 12)
	fmt.Println(n)
}

func HashFunction(word string, buckets int) int{
	letter := int(word[0])
	bucket := letter % buckets
	return bucket
}
