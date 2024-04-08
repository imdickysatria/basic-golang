package main

import "fmt"

func main() {
	a := printResult("saya sedang")
	fmt.Println(a)

	result := add(10,29)
	fmt.Println(result)

	fmt.Println("=============")

	luas, keliling := calculate(10,2)
	fmt.Println(luas)
	fmt.Println(keliling)

	fmt.Println("=============")
	kali,kurang := timesMin(10,2)
	fmt.Println(kali)
	fmt.Println(kurang)
}

func printResult(a string) string {
	newA := a + " belajar golang"

	return newA
}

func add(b, c int) int  {
	result := b + c
	return result
	
}
func calculate(panjang, lebar int) (int,int){
	luas  := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

func timesMin(g, h int) (kali,kurang int){
	kali = g * h
	kurang = g-h
	return
}