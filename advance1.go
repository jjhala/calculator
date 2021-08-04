package main

import (
	"fmt"
	"os"
	"strconv"
)

func main1() {
  
  // This takes 3 parameters and performs the operation. For example - go run main.go 2 + 5 
  //Also throws an error if there are less than 3 arguments. 
  
  // this does not work if I specify 2 * 5 because it takes * as something else. But if I do 2 "*" 5 then it performs the multiplication. 

	if len(os.Args) <= 3 {
		fmt.Println("Need 3 arguments please!")
		os.Exit(1)
	}
	a, _ := strconv.Atoi(os.Args[1]) //using strconv package to convert the string from arguments to integer
	operation := os.Args[2]
	b, _ := strconv.Atoi(os.Args[3])
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
