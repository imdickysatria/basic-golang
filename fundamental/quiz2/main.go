package main

import (
	"errors"
	"fmt"
)

func main() {
	scores := []int{10, 5, 8, 9, 7}
	total := sum(scores)
	fmt.Println(total)

	fmt.Println("===========")

	result, err := calculation(10,2,"/")
	if err != nil {
		fmt.Println("terjadi kesalahan")
		fmt.Println(err.Error())
	}
	fmt.Println(result)
	// result, err := calculation(10,2,"-")
	// result, err := calculation(10,2,"*")
	// result, err := calculation(10,2,"/")
	// result, err := calculation(10,2,"=")
}

func sum(numbers []int) int {
	var total int

	for _, number := range numbers {
		total = total + number
	}
	return total
}

func calculation(a,b int,operator string ) (int,error)  {
	var result int
	var errorResult error
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "*":
		result = a * b
	case "/":
		result = a / b
	default:
		errorResult = errors.New("Unknown Operator")
	}

	return result, errorResult
	
}