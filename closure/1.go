package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	increment := wrapper()
	fmt.Printf("%v\n", increment())
	fmt.Printf("%v\n", increment())
	fmt.Printf("%v\n", increment())
	fmt.Printf("%v\n", increment())
	fmt.Printf("%v\n", increment())
}
