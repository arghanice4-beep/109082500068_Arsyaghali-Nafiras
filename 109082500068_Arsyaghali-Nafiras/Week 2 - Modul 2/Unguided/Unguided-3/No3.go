package main

import "fmt"

func main() {
	var beratTotalGram int
	fmt.Print("Masukkan total berat (gram): ")
	fmt.Scanf("%d", &beratTotalGram)
	fmt.Println("===== Detail Perhitungan =====")
	const HargaPerKg = 10000
	kg := beratTotalGram / 1000
	sisaGram := beratTotalGram % 1000
	biayaKg := kg * HargaPerKg
	var biayaGram int
	if kg > 10 {
		biayaGram = 0
	} else {
		if sisaGram >= 500 {
			biayaGram = sisaGram * 5
		} else {
			biayaGram = sisaGram * 15
		}
	}

	totalBiaya := biayaKg + biayaGram

	fmt.Printf("Detail Berat: %d kg + %d gr\n", kg, sisaGram)
	fmt.Printf("Detail Biaya: Rp. %d (kg) + Rp. %d (gr)\n", biayaKg, biayaGram)
	fmt.Printf("Total Biaya : Rp. %d\n", totalBiaya)
	fmt.Println("---------------------------------------")
}
