package main

import "fmt"

func main() {
	var target, tabungan, total, hari int

	fmt.Print("Masukkan target tabungan: ")
	fmt.Scanln(&target)
	total = 0
	hari = 0

	for total < target {
		hari++
		fmt.Printf("Masukkan jumlah tabungan hari ke-%d: ", hari)
		fmt.Scanln(&tabungan)
		total = total + tabungan
	}

	fmt.Printf("selamat! target tercapai dalam %d hari\n", hari)
	fmt.Printf("Total tabungan anda terkumpul: %d\n", total)
}