package main

import "fmt"

func main() {

	var (
		a         int
		b         int
		operation string
	)
	a = 2
	b = 14
	operation = "*" //change this value to perform different operations
	switch operation {
	case "+":
		fmt.Println("Result is", a+b)
	case "-":
		fmt.Println("Result is", a-b)
	case "*":
		fmt.Println("Result is", a*b)
	case "/":
		fmt.Println("Result is", a/b)
	default:
		fmt.Println("Not a valid operation")
	}
}
