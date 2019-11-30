package main

import (
	"fmt"
)

type Person struct{
	First string
	last string
	age int
}

type Company struct {
	Person
	First string
	strength int
}

func  main()  {
	P1 := Company{
		Person : Person{
			First: "Digambar",
			last: "Patil",
			age: 33,
		},
		First: "NVDIA",
		strength: 10000,
	}
	fmt.Println(P1.First, P1.strength, P1.Person.First)
}