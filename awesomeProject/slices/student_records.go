package main

import (
	"fmt"
)

func records() {
	records := make([][]string, 0)

	//student1
	student1 := make([]string, 4)
	student1[0] = "Atul"
	student1[1] = "Pune"
	student1[2] = "100.00"
	student1[3] = "80.00"

	records = append(records, student1)

	//student2
	student2 := make([]string, 4)
	student2[0] = "Sham"
	student2[1] = "Kolhapur"
	student2[2] = "90.00"
	student2[3] = "85.00"

	records = append(records, student2)

	fmt.Println("Student 1 record - ", student1)
	fmt.Println("Student 2 record - ", student2)
	fmt.Println("All class records - ", records)
}
