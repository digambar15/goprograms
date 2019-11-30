package main

import (
	"fmt"
	"runtime"
)

var x int
var y float64
var a int8 = -128
//var b int8 = -129
func main(){
	x = 42
	y = 40.123
	fmt.Println(x)
	fmt.Println(y)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)

	fmt.Printf("%v\n", a)

	//assigned to int values ranges from -128 to 128
	//fmt.Println(b)

	fmt.Println(runtime.GOOS)
	fmt.Println("Go Architecture ", runtime.GOARCH)
}