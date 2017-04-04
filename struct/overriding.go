package main

import "fmt"

type person struct {
	First string
	Last  string
	Age   int
}

type doubleZero struct {
	person
	First         string
	LicenseToKill bool
}

func main() {
	p1 := doubleZero{
		person: person{
			First: "James",
			Last:  "Bond",
			Age:   20,
		},
		First:         "Double Zero Seven",
		LicenseToKill: true,
	}

	p2 := doubleZero{
		person: person{
			First: "Miss",
			Last:  "Shang",
			Age:   123,
		},
		First:         "xiejging",
		LicenseToKill: false,
	}

	fmt.Printf("%v", p1)
	fmt.Printf("%v", p2)
}
