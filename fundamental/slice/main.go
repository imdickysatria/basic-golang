package main

import "fmt"

func main (){
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "PS1")
	gamingConsoles = append(gamingConsoles, "PS2")
	gamingConsoles = append(gamingConsoles, "PS3")
	gamingConsoles = append(gamingConsoles, "PS4")

	gamingConsoles2 := []string{"Xbox 1", "Xbox 2", "Xbox 3", "Xbox 4"}



	// fmt.Println(gamingConsoles)
	// fmt.Println(gamingConsoles2)


	for _, console := range gamingConsoles{
		fmt.Println(console)
	}

	for _, console2 := range gamingConsoles2{
		fmt.Println(console2)
	}
}