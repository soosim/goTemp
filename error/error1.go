package main

import (
	"errors"
	"fmt"
)

type Error1 struct {
	arg    int
	errMsg string
}

func (e *Error1) Error() string {
	return fmt.Sprintf("%s", e.errMsg)
}

type Error2 struct {
	arg    string
	errMsg string
}

func (e *Error2) Error() string {
	return fmt.Sprintf("%s", e.errMsg)
}

func test0() error {
	return errors.New("errors.New() - test0()")
}

func test1() error {
	return fmt.Errorf("fmt.Errorf() - test1()")
}

func test2(arg int) *Error1 {
	return &Error1{arg, "Error1{} - text2()"}
}

func test3(arg string) error {
	return &Error2{arg, "Error2{} test3()"}
}

func test4(arg string) *Error2 {
	return &Error2{arg, "Error2{} - test4()"}
}

func errType(err interface{}) {
	switch e := err.(type) {
	case *Error1:
		fmt.Println("Type:Error1 ", e)
	case *Error2:
		fmt.Println("Type:Error2 ", e)
	case error:
		fmt.Println("Type:error ", e)
	default:
		fmt.Println("Type default")
	}
}

func main() {
	errType(test0())
	errType(test1())
	errType(test2(2))
	errType(test3("asdafasf"))
	errType(test4("safsajf"))
}
