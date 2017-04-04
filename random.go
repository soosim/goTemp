package main

import (
	"fmt"
	"math/rand"
	"time"
)

type MathFunc func(int, int) int

func main() {
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	fmt.Println("----------------------")

	rand.Seed(time.Now().UnixNano())
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))
	fmt.Println(rand.Intn(100))

	fmt.Println("----------------------")

	fmt.Println("Random Int:", rand.Int())
	fmt.Println("Random Float:", rand.Float32())
	fmt.Println("Random Permutation of N ints:", rand.Perm(8))
	fmt.Println("Normalized:", rand.NormFloat64())
}
