package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func switchOnType(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	switchOnType(7)
	switchOnType("asfasg")
	var t = contact{"asdsagasf", "asd"}
	switchOnType(t)
	switchOnType(t.greeting)
	switchOnType(t.name)
}

func multiple() {
	switch "Jenny" {
	case "Time", "Jenny":
		fmt.Println("asdasfagasf")
	case "Marche", "Decome":
		fmt.Println("Decome or Narchhe")
	default:
		fmt.Println("Default values")
	}
}

func fallThrough() {
	switch "T" {
	case "T":
		fmt.Println("T")
		fallthrough
	case "U":
		fmt.Println("U")
	case "T":
		fmt.Println(".........")
		fallthrough
	}
}
