package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "Shang",
			Age:   20,
		},
		LicenseToKill: false,
	}
	fmt.Printf("%v", p1)
	fmt.Printf("%v", p2)
}
