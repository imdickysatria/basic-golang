package main

import "fmt"
func main(){
	const tinggi = 100
	fmt.Println("nilai tingi =",tinggi)

	var jumlah = ((2+5)*10)/2

	var uji = jumlah !=20

	fmt.Println("hasil dari penjumlahan =",jumlah)
	fmt.Println("Test uji =",uji)

	var kiri = true
	var kanan = false

	hasil := kiri && kanan
	hasilOr := kiri || kanan
	notKiri := !kiri

	fmt.Println("hasil and =",hasil)
	fmt.Println("hasil or =",hasilOr)
	fmt.Println("hasil notKiri =", notKiri)

}
