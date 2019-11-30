package main

import (
	"fmt"
	"math"
	"sort"
)

func main(){
	s := []float64{5.2, -1.3, 0.7, -3.8, 2.6}
	sort.Float64s(s)
	fmt.Println(s)

	s = []float64{math.Inf(1), math.NaN(), math.Inf(-1), 0.0} // unsorted
	sort.Float64s(s)
	fmt.Println(s)

	ss := []float64{0.7, 1.3, 2.6, 3.8, 5.2} // sorted ascending
	fmt.Println("Sorted in ascending order : ", sort.Float64sAreSorted(ss))

	ss = []float64{5.2, 3.8, 2.6, 1.3, 0.7} // sorted descending
	fmt.Println("Sorted in descending order : ", sort.Float64sAreSorted(ss))

	ss = []float64{5.2, 1.3, 0.7, 3.8, 2.6} // unsorted
	fmt.Println("Not sorted : ",sort.Float64sAreSorted(ss))

}
