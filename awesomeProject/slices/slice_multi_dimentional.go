package main

import "fmt"

func main(){
	jb := []string{"James", "Bond", "chocolate", "martini"}
	fmt.Println(jb)

	mp := []string{"Miss", "Patil", "Strawberry", "sweet"}
	fmt.Println(mp)

	mt := [][]string{jb, mp}
	fmt.Println(mt)

	for i, v := range mt{
		fmt.Println(i, " - ", v)
	}

	transactions := make([][]int, 0, 3)
	fmt.Println("length of transactions is", len(transactions))
	fmt.Println("Capacity of transactions are ", cap(transactions))
	for i := 0; i < 4; i++{
		transaction := make([]int, 0)
		for j := 0; j < 2; j++{
			transaction = append(transaction, j)
		}
		transactions = append(transactions, transaction)
	}
	fmt.Println("Transactions are ", transactions)
	fmt.Println("length of transactions is", len(transactions))
	fmt.Println("Capacity of transactions are ", cap(transactions))

	myslice := make([]int, 2)
	fmt.Println(myslice)

	myslice[1] = 7
	fmt.Println(myslice)

	myslice[1]++
	fmt.Println(myslice)
}
