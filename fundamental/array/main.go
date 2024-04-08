package main

import "fmt"

func main()  {

	// var languages [5]string
	// languages[0] = "GO"
	// languages[1] = "Python"
	// languages[2] = "PHP"
	// languages[3] = "Java"
	// languages[4] = "Javascript"

	languages := [...]string{
		"go",
		"ruby",
		"python",
		"c++",
		"php",
		"js",
	}
	// fmt.Println(languages)

	// length := len(languages)
	// fmt.Println(length)

	for index, lang := range languages{
		fmt.Println("index :", index, "language :", lang)
	}
}