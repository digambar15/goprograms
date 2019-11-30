package main

import (
	"fmt"
)

type Emp struct {
	first string
	last string
	age int
}

type Comp struct {
	Emp
	first string
	strength int
}

func (p Emp) greeting(){
	fmt.Println("Hello James")
}

func (c Comp) greeting(){
	fmt.Println("Welcome to NVDIA")
}

func  main()  {
	P1 := Comp{
		Emp : Emp{
			first: "Digambar",
			last: "Patil",
			age: 33,
		},
		first: "NVDIA",
		strength: 10000,
	}
	fmt.Println(P1.first, P1.strength, P1.Emp.first)
	P1.greeting()
	P1.Emp.greeting()
}
