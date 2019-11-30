package main

import "fmt"

var i int
var j string
var k bool
type home int
var a home
var b int
func main(){
	x := 42
	y := "James, Bond"
	z := true

	fmt.Println("Values of x, y and z are ", x, y, z)
	fmt.Printf("Value of x is %v\n", x)
	fmt.Printf("Value of y is %v\n", y)
	fmt.Printf("Value of z is %v\n", z)

	fmt.Println("Zero values of i, j and k variables are ", i, j, k)

	//Use Sprintf to assign the values to one single string
	s := fmt.Sprintf("%v\t %v\t %v\t", x, y, z)
	fmt.Println(s)

	//Create your own variable type and assign a value to it and print it
	fmt.Println("Value of a is ", a)
	fmt.Printf("Type of variable a is %T", a)
	var a = 50
	fmt.Println("Value of a is ", a)

	//Conversion
	b = int(a)
	fmt.Println("Value of b after conversion is ", b)
	fmt.Printf("Type of variable b is %T", b)
	}