package main

import (
	"fmt"
	"fundamental/calculation"
)

func main()  {
	fmt.Println("haloo")

	a := 9
	b := 9
	result := calculation.Add(a, b)
	// result2 := calculation.Min(a, b)
	resul3 := calculation.Times(a, b)
	fmt.Println(result)
	// fmt.Println(result2)
	fmt.Println(resul3)

	var name string = "riyan"

	fmt.Println(name)
}