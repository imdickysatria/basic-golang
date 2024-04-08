package main

import "fmt"

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["ruby"] = 9
	// myMap["go"] = 9
	// myMap["java"] = 8
	// myMap["php"] = 10

	// fmt.Println(myMap["ruby"])

	myMap := map[string]string{
	"ruby": "is beautiful", 
	"go": "is super fast",
	"java" : "hype",
	}

	for key, value := range myMap{
		fmt.Println("key :", key, "value :", value)
	}

	fmt.Println("===============")

	delete(myMap, "ruby")
	
	for key, value := range myMap{
		fmt.Println("key :", key, "value :", value)
	}

	fmt.Println("===============")

	value, isAvailable := myMap["python"]
	fmt.Println(isAvailable)
	fmt.Println(value)
}