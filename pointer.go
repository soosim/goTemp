package main

import "fmt"

type Dog struct {
	Name  string
	Color string
}

func main() {
	Spot := Dog{Name: "Spot", Color: "brown"}

	SpotPonter := &Spot

	SpotPonter.Color = "black"

	fmt.Println(Spot)
}
