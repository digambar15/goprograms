package main

import (
	"fmt"
)

func transactions() {
	transactions := make([][]int, 0, 3)

	for i := 0; i < 3; i++ {
		transaction := make([]int, 3)
		for j := 0; j < 3; j++ {
			transaction[j] = j
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println("Transaction set - ", transactions)
}
