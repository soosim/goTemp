package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func mian() {
	var guess int
	var count int

	num := rand.Int(100) + 1

	fmt.Println(">> between -100 <<")

	for {
		fmt.Print("Guess: ")
		_, err := fmt.Scanf("%d", &guess)
		if err == nil {
			count += 1
			if guess > num {
				fmt.Println("Too HIGH")
			} else if guess < num {
				fmt.Println("Too LOW")
			} else {
				fmt.Println("Good")
			}
		} else {
			fmt.Println(">> Please input a number")
		}
	}
}
