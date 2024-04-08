package main

import "fmt"

func main() {
	students := []map[string]string{
		{
			"name":  "agung",
			"score": "A",
		},
		{
			"name":  "supri",
			"score": "B",
		},
		{
			"name":  "adi",
			"score": "A",
		},
		{
			"name":  "mario",
			"score": "c",
		},
	}

	fmt.Println(students)

	for _, student := range students{
		fmt.Println(student["name"],"------", student["score"])
	}

}