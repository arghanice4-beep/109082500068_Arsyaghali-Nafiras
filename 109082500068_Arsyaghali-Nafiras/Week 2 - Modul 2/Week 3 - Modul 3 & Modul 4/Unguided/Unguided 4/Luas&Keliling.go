package main

import (
	"fmt"
	"math"
)

func hitungPersegi(sisi int) {
	luas := sisi * sisi
	keliling := 4 * sisi
	fmt.Printf("--- Hasil Persegi ---\n")
	fmt.Printf("Luas: %d\n", luas)
	fmt.Printf("Keliling: %d\n\n", keliling)
}

func hitungPersegiPanjang(panjang, lebar int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)
	fmt.Printf("--- Hasil Persegi Panjang ---\n")
	fmt.Printf("Luas: %d\n", luas)
	fmt.Printf("Keliling: %d\n\n", keliling)
}

func hitungLingkaran(jarijari float64) {
	luas := math.Pi * jarijari * jarijari
	keliling := 2 * math.Pi * jarijari
	fmt.Printf("--- Hasil Lingkaran ---\n")
	fmt.Printf("Luas: %.2f\n", luas)
	fmt.Printf("Keliling: %.2f\n\n", keliling)
}

func main() {
	var pilihan int

	fmt.Println("--- PROGRAM BANGUN DATAR ---")
	fmt.Println("1. Hitung luas & keliling persegi")
	fmt.Println("2. Hitung luas & keliling persegi panjang")
	fmt.Println("3. Hitung luas & keliling lingkaran")
	fmt.Print("Pilihan : ")
	fmt.Scan(&pilihan)

	switch pilihan {
	case 1:
		var s int
		fmt.Print("Masukkan sisi: ")
		fmt.Scan(&s)
		hitungPersegi(s)
	case 2:
		var p, l int
		fmt.Print("Masukkan panjang: ")
		fmt.Scan(&p)
		fmt.Print("Masukkan lebar: ")
		fmt.Scan(&l)
		hitungPersegiPanjang(p, l)
	case 3:
		var r float64
		fmt.Print("Masukkan jari-jari: ")
		fmt.Scan(&r)
		hitungLingkaran(r)
	default:
		fmt.Println("Pilihan tidak valid.")
	}
}