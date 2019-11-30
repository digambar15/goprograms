package main

import "fmt"

func main(){
	fmt.Println(company("NVDIA", 30))
}

func company(name string, age int) (string, int){
	name = name
	age = age + 10
	return name, age
}
