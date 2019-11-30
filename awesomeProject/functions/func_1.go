package main

import "fmt"

func main(){
	greet("Welcome John!")
	greet("Wow you finally Here. Jane!")

	person("Digambar", "Patil", 33)
}

func greet(name string){
	fmt.Println(name)
}

func person(first ,last string, age int){
	fmt.Println(first, last, age)
}
