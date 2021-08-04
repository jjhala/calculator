package main

import (
	"fmt"
)

func main() {

	a := 2
	b := 14
	operation := "*"
	doMath(operation, a, b)
}

// Sending the a and b parameters as interface type so it can handle both int and float.

func doMath(operation string, a, b interface{}) {
	switch v1 := a.(type) {
	case int:
		switch operation {
		case "+":
			fmt.Println("Result is", v1+b.(int))
		case "-":
			fmt.Println("Result is", v1-b.(int))
		case "*":
			fmt.Println("Result is", v1*b.(int))
		case "/":
			fmt.Println("Result is", v1/b.(int))
		default:
			fmt.Println("Not a valid operation")
		}
	case float64:
		switch operation {
		case "+":
			fmt.Println("Result is", v1+b.(float64))
		case "-":
			fmt.Println("Result is", v1-b.(float64))
		case "*":
			fmt.Println("Result is", v1*b.(float64))
		case "/":
			fmt.Println("Result is", v1/b.(float64))
		default:
			fmt.Println("Not a valid operation")
		}
	default:
		fmt.Println("Has to be either integer or float type")
	}
}
