package main

import "fmt"

type person struct {
	Name string
	Age  int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func (p person) Greeting() {
	fmt.Println("I'm just a regular person")
}

func (dz doubleZero) Greeting() {
	fmt.Println("Miss Moneypenny, so good to see you")
}

func main() {
	p1 := person{
		Name: "I am Flemming",
		Age:  44,
	}

	p2 := doubleZero{
		person: person{
			Name: "Jame Bond",
			Age:  23,
		},
		LicenseToKill: false,
	}

	p1.Greeting()
	p2.Greeting()
	p2.person.Greeting()
}
