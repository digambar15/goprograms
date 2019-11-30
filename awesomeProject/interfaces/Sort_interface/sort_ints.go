package main

import (
	"fmt"
	"sort"
)

func main(){
	s := []int{5,2,4,3,6,1}
	sort.Ints(s)
	fmt.Println(s)

	s = []int{1,2,3,4,5,50}
	fmt.Println("sorted in ascending order : ", sort.IntsAreSorted(s))

	s = []int{5,4,3,2,1,0}
	fmt.Println("Sorted in descending order : ", sort.IntsAreSorted(s))

	s = []int{5,3,4,1,2}
	fmt.Println("Not sorted : ", sort.IntsAreSorted(s))

	i := sort.Reverse(sort.IntSlice(s))
	fmt.Printf("Type of Reverse %T \n", i)
	sort.Sort(i)
	fmt.Println(s)
}
