package main

import "fmt"

var y = 50
var z = `James said "Welcome to Go"`
func main(){
	n, _ := fmt.Print("Hello World ", 40, true)
	fmt.Println(n)

	x := 40
	fmt.Println("X = ", x)

	fmt.Println("Y = ", y)

	fmt.Println("Z = ", z)
	fmt.Printf("%T\n", z)

}