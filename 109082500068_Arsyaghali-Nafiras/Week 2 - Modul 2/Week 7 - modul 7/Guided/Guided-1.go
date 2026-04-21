package main

import "fmt"

type desimal float64

func main() {
	var alas, tinggi, luas desimal

	fmt.Print("Masukkan alas segitiga: ")
	fmt.Scanln(&alas)

	fmt.Print("Masukkan tinggi segitiga: ")
	fmt.Scanln(&tinggi)

	luas = desimal(0.5) * alas * tinggi

	fmt.Printf("Luas segitiga dengan alas %.2f dan tinggi %.2f adalah: %.2f\n", alas, tinggi, luas)
}