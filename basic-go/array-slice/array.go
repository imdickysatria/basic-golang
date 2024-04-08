package main

import "fmt"

func Array()  {
	var buah [8]string
	buah[0] = "apel"
	buah[1] = "per"
	buah[2] = "timun"
	buah[3] = "cabe"
	buah[4] = "melon"
	buah[5] = "nangka"
	buah[6] = "jambu"
	buah[7] = "jeruk"


	fmt.Println("jumlah data :", len(buah))
	fmt.Println("data array: ",buah)

	for i:=0;i < len(buah);i++{
		fmt.Println("buah apa aja :", buah[i])
	}

	var angka = [...]int{1,2,4,5,6,9}

	fmt.Println("jumlah data:", len(angka))
	fmt.Println("data array :", angka)

	var fruit []string
	fruit = append(fruit, "semangka")
	fruit = append(fruit, "timun")
	fruit = append(fruit, "apel")
	fruit = append(fruit, "keju")

	fmt.Println("jumlah data : ", len(fruit))
	fmt.Println("jumlah data : ", fruit)

	fruit = append(fruit, "blimbing")
	fruit = append(fruit, "duku")

	fmt.Println("jumlah data baru :", len(fruit))
	fmt.Println("total data baru :", fruit)

	for g:=0; g<len(fruit);g++{
		fmt.Println("ini adalah buah :",fruit[g]) 
	}
	
}