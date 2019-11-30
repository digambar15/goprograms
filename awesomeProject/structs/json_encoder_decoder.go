package main

import (
	"encoding/json"
	"fmt"
)

type PS struct {
	name 	string
	address string	`json:"-"`
	pin 	int		`json:"wisdom score"`
}

func main()  {
	js := PS{"Digambar", "Gavhan", 100008}
	fmt.Println(js)
	bs, _ := json.Marshal(js)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}