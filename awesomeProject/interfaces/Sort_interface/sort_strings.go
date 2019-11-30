package main

import (
	"fmt"
	"sort"
)

func main(){
	s := []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Strings(s)
	fmt.Println(s)

	s = []string{"Alpha", "Bravo", "Delta", "Go", "Gopher", "Grin"}
	fmt.Println("Sorted in ascending order ", sort.StringsAreSorted(s))
	s = []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	fmt.Println("Not sorted", sort.StringsAreSorted(s))

	sort.StringSlice(s).Sort()  //Using StringSlice interface
	fmt.Println(s)
	fmt.Println("Go found in string list", sort.StringSlice(s).Search("Go"))

	s = []string{"Go", "Bravo", "Gopher", "Alpha", "Grin", "Delta"}
	sort.Sort(sort.StringSlice(s))
	fmt.Println("Strings are sorted ", s)

	fmt.Printf("%T string \n", sort.StringSlice(s))
	i := sort.Reverse(sort.StringSlice(s))
	fmt.Printf("Type of reverse is %T \n", i)
	sort.Sort(i)
	fmt.Println("Sorted in reverse order ", s)
}