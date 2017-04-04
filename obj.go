package main

import (
	"fmt"
)

type Dog struct {
	Name  string
	Color string
}

func (d Dog) Call() {
	fmt.Printf("Here comes a %s dog, %s.\n", d.Color, d.Name)
}

func main() {
	Spot := Dog{Name: "Spot", Color: "brown"}
	Spot.Call()

	var Rover Dog
	Rover.Name = "Rover"
	Rover.Color = "light blondish with some dark and very.."
	Rover.Call()
}
