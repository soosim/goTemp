package main

import (
	"fmt"
	"time"
)

func printSlowly(s string, n int) {
	for i := 0; i < n; i++ {
		time.Sleep(300 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	printSlowly("directly functioning", 3)

	printSlowly("red fish goroutine", 3)
	printSlowly("blue fish goroutine", 3)

	go func(ss string, nn int) {
		for i := 0; i < nn; i++ {
			fmt.Println(i, ss)
			time.Sleep(150 * time.Millisecond)
		}
	}("annony fish goroutine", 3)

	var input string
	fmt.Scanln(&input)
	fmt.Println("Done.")

	fmt.Println("End")
}
