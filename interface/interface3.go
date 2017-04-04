package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	redius float64
}

type square struct {
	side float64
}

func (c circle) area() float64 {
	return math.Pi * c.redius * c.redius
}

func (s square) area() float64 {
	return s.side * s.side
}

func info(z shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}

func main() {
	s := square{100}
	c := circle{5}

	info(s)
	info(c)
	fmt.Println("Totla Area: ", totalArea(c, s))
}
