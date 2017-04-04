package main

import "fmt"

func main() {
	var empty []int

	alphas := []string{"abc", "def", "ghi", "jki"}

	empty = append(empty, 123)
	empty = append(empty, 456)

	fmt.Printf("%v \n", empty)

	alphas = append(alphas, "pqr", "stu")
	fmt.Printf("%v \n", alphas)

	fmt.Println("Length:", len(alphas))

	fmt.Println(alphas[1])

	alpha2 := alphas[1:3]
	fmt.Println(alpha2)

	if elemExists("def", alphas) {
		fmt.Println("Exists!")
	}
}

func elemExists(s string, a []string) bool {
	for _, v := range a {
		if v == s {
			return true
		}
	}
	return false
}
