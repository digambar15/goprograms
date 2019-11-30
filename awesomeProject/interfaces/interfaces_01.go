package main

import (
	"fmt"
	"math"
)

type square struct {
	side float64
}

func (s square) area() float64{
	return s.side * s.side
}

type circle struct {
	radius float64
}

func (c circle) area() float64{
	return math.Pi * c.radius * c.radius
}

type Shape interface {
	area() float64
}

func info(z Shape){
	fmt.Println(z)
	fmt.Println(z.area())
}

func main()  {
	s := square{10}
	c := circle{20}
	info(s)
	info(c)
}