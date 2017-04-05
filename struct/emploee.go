package main

import "fmt"

type Employee struct {
	ID        int
	Name      string
	Address   string
	Dob       string
	Position  string
	Salary    int
	ManagerID int
}

var dilbert Employee

func main() {
	dilbert.Salary = 5000
	dilbert.Position = "sdasfasfasgasf"

	position := &dilbert.Position
	*position = "xiejinlong " + *position
	fmt.Println(dilbert)
}
