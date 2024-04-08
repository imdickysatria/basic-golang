package main

import "fmt"

func main()  {
	var nilai int
	fmt.Println("masukkan nilai : ")

	fmt.Scan(&nilai)
	if nilai == 10 {
		fmt.Println("sangat bagus")
	} else if nilai <10  && nilai >=5{
		fmt.Println("bagus")
	}else if nilai <5 && nilai > 1 {
		fmt.Println("Buruk")
	}else{
		fmt.Println("Coba lagi")
	}

	var nlai2 int
	fmt.Println("masukkan nilai: ")
	fmt.Scan(&nlai2)

	switch nlai2{
	case 10 :
		fmt.Println("sangat bagus")
	case 8:
		fmt.Println("bagus")
	case 5:
		fmt.Println("buruk")
	default:
		fmt.Println("coba lagi")
	}
}