package main

import "fmt"

type home struct {
	address string
	pin int
}

func (h home) findWay()(string, int){
	return h.address, h.pin
}

func main(){
	h := &home{
		address: "Gavhan, Tasgaon",
		pin: 400000,
	}
	fmt.Println(h)
	fmt.Printf("%T\n", h)
	k, v := h.findWay()
	fmt.Println(k, v)
}
