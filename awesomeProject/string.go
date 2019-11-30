package main

import "fmt"

func main(){
	s := "Hello, World"
	fmt.Println(s)
	s = "Digambar Patil"
	fmt.Println(s)
	b := []byte(s)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}
