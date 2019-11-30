package main

import "fmt"

type PersonData struct {
	first string
	last  string
	age   int
}

type DoubleData struct {
	PersonData
	first string
	last  string
}

func (p PersonData) fullname() string {
	return p.first + p.last
}

func (p PersonData) allData() PersonData {
	return p
}

func main() {
	st := PersonData{"Digambar", "Patil", 64}
	p1 := PersonData{"Noopur", "Landge", 27}
	p2 := DoubleData{
		PersonData: PersonData{
			first: "Girish",
			last:  "Landge",
			age:   33,
		},
		first: "Amey",
		last:  "Raynade",
	}
	fmt.Println(st.fullname())
	fmt.Println(st)
	fmt.Println(p1.allData())
	fmt.Println(p2)
	fmt.Println(p2.first, p2.age)
	fmt.Println(p2.PersonData.first, p2.PersonData.age)

}
