package main

import (
	"fmt"
	"regexp"
)

func basic_regexes() {
	pattern := "[0-9]+"
	str := "The 12 monkeys ate 48 bananas"

	re, err := regexp.Compile(pattern)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// 1.test compiled pattern matches string
	if re.MatchString(str) {
		fmt.Println("Yes ,matched a number")
	} else {
		fmt.Print("No,no match")
	}

	// 2.return first match
	result_two := re.FindString(str)
	fmt.Println("Number matched:", result_two)

	// 3.return n matched, use -1 to  find all matched
	result_three := re.FindAllString(str, 2)
	for i, v := range result_three {
		fmt.Printf("Match %d:%s\n", i, v)
	}
}

func case_insensitive() {
	ptn := `(?i)^t`
	str := "To be or not to be"

	re, err := regexp.Compile(ptn)
	if err != nil {
		fmt.Println("Error compiling regex", err)
	}

	// match string
	result := re.FindString(str)
	fmt.Println("Result:", result)
}

func sub_matches() {
	str := "Hello @world@ Match"
	sub_re, _ := regexp.Compile("@(.*)@")
	m := sub_re.FindStringSubmatch(str)
	if len(m) > 1 {
		fmt.Println(m[1])
	}

	str2 := "A [word] here and [there]"
	esc_pattern := `\[(.*?)\]\`
	esc_re, _ := regexp.Compile(esc_pattern)
	fmt.Println(esc_re.FindStringSubmatch(str2))

	fmt.Println(esc_re.FindAllStringSubmatch(str2, -1))
}

func main() {
	basic_regexes()
	case_insensitive()
	sub_matches()
}
