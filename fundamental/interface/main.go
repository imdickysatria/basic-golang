package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar

}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 8, Lebar: 4}
	luas := seberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang: ", luas)

	persegi := Persegi{Sisi: 8}
	luas = seberapaLuas(persegi)
	fmt.Println("Luas Persegi: ", luas)
}

func seberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}