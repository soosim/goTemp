package main

import (
	"encoding/json"
	"fmt"
)

type person struct {
	First string
	Last  string
	Age   int
}

func main() {
	p1 := person{"James", "Bond", 25}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
	fmt.Println(string(bs))
}
