package main

import "fmt"

func greeting() func() string{
	return func() string{
		return "Hello World"
	}
}

func main(){
	greet := greeting()
	fmt.Println(greet())
}
