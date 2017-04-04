package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	p1 := &person{"Jame", 20}
	fmt.Printf("%T\n", p1)
	fmt.Println(p1.name)
	fmt.Println(p1.age)
}
