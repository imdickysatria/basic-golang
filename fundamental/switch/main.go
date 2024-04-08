package main

import (
	"fmt"
)

func main()  {

	number := 5
	switch number {
	case 1:
		fmt.Println("satu")
	case 2:
		fmt.Println("dua")
	case 3:
		fmt.Println("tiga")
	default:
		fmt.Println("coba lagi")
	}

	// for i := 1; i <= 10; i++{
	// 	fmt.Println("halo riyan ganteng")
	// } 
	
	// i := 1
	// for i < 100{
	// 	fmt.Println("halo riyan ganteng", i)
	// 	i++
	// } 

	// title := "Riyan Dicky Satria"

	// for _, letter := range title{
	// 	fmt.Println("letter :", string(letter))
	// }

	title2 := "Golang the best language"

	for index, test := range title2{
		if index % 2 == 0 {
			fmt.Println("index :",index, "letter :", string(test))
		}
	}

	for index, test := range title2{
		letterString := string(test)
		if letterString == "a" || letterString == "i" || letterString == "u" || letterString== "e" || letterString == "o" {
			fmt.Println("index :", index, "letter :", string(test))
		}
	}
}