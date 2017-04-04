package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Person struct {
	Name, City string
}

var person Person

var people []Person

func main() {
	json_str := "{\"Name\":\"Xiejinlong\",\"sex\":\"male\",\"city\":\"shiyan\"}"

	if err := json.Unmarshal([]byte(json_str), &person); err != nil {
		fmt.Println("Error paring JSON:", err)
	}

	fmt.Printf("Name: %v, City:%v\n", person.Name, person.City)

	file, err := ioutil.ReadFile("./1.json")
	if err != nil {
		fmt.Println("error reading file")
	}

	if err := json.Unmarshal(file, &people); err != nil {
		fmt.Println("Error parsing JSON", err)
	}

	fmt.Println(people)

	// json, err := json.Marshal(people)
	if err != nil {
		fmt.Println("JSON Encoding Error", err)
	}
	for _, v := range people {
		fmt.Printf("%s => %s\n", v.Name, v.City)
	}
	// fmt.Println(string(json))

}
