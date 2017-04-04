package main

import "fmt"

func main() {
	num := 1
	if num > 3 {
		fmt.Println("Many")
	} else {
		fmt.Println("Low")
	}

	if num == 1 {
		fmt.Println("one")
	} else if num == 2 {
		fmt.Println("two")
	} else {
		fmt.Println("Many")
	}

	switch {
	case num == 1:
		fmt.Println("one")
	case num == 2:
		fmt.Println("two")
	case num > 2:
		fmt.Println("many")
	default:
		fmt.Println("Throw over boat")
	}
}
