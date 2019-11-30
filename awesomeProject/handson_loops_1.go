package main

import "fmt"

func main(){
	for i:=0; i<100; i++{
		fmt.Println(i)
	}

	for i := 65; i <=70; i++{
		for j:=0; j<3; j++{
			fmt.Printf("%#U\n", i)
		}
	}

	k := 1985
	for k > 2019 {
		fmt.Println(k)
		k++
	}
	fmt.Println("Total age is ", 2019-1985)
}
