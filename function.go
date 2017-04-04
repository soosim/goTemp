package main

import (
	"fmt"
	"math"
)

func Echo(s string) {
	fmt.Println(s)
}

func Say(s string) string {
	pharase := "hello " + s
	return pharase
}

func Say2(s string) (phrase string) {
	phrase = "Hello " + s
	return
}

func Divide(x, y float64) (float64, float64) {
	q := math.Trunc(x / y)
	r := math.Mod(x, y)
	return q, r
}

func Divide2(x, y float64) (q, r float64) {
	q = math.Trunc(x / y)
	r = math.Mod(x, y)
	return
}

func main() {
	x, y := Divide2(1332, 90)
	fmt.Println(x, y)
}
