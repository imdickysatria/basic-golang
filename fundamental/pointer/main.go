package main

import "fmt"

func main() {
	// numberA := 5
	// numberB := &numberA

	// fmt.Println(numberA)
	// fmt.Println(numberB)
	// fmt.Println(*numberB)


	// *numberB = 10

	// fmt.Println(*numberB)
	// fmt.Println(numberA)

	// var numberA int = 5
	// var numberB *int = &numberA

	// fmt.Println(numberA)
	// fmt.Println(*numberB)

	// numberA = 20

	// fmt.Println(numberA)
	// fmt.Println(*numberB)

	number := 5 
	fmt.Println("nilai awal : ",number)
	change(&number, 100)

	fmt.Println("nilai akhir : ",number)
}
func change(old *int, new int)  {
	*old = new 
	fmt.Println("alamat memori ",&old)
	fmt.Println("di dalam function ",old)
}

