package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type square struct {
	side float64
}

type circle struct {
	redius float64
}

func (s square) area() float64 {
	return s.side * s.side
}

func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	s := square{10}
	c := circle{10}
	info(s)
	info(c)
}
