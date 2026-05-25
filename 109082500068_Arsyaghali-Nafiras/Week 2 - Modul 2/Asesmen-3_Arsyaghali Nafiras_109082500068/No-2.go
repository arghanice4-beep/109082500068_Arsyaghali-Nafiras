package main

import "fmt"

type Pemain struct {
	NamaDepan    string
	NamaBelakang string
	Gol          int
	Assist       int
}

func main() {
	var n int
	
	fmt.Println("Masukkan Data Input :")
	fmt.Scan(&n)

	daftarPemain := make([]Pemain, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&daftarPemain[i].NamaDepan, &daftarPemain[i].NamaBelakang, &daftarPemain[i].Gol, &daftarPemain[i].Assist)
	}

	for i := 1; i < n; i++ {
		kunci := daftarPemain[i]
		j := i - 1
		for j >= 0 && (daftarPemain[j].Gol < kunci.Gol || (daftarPemain[j].Gol == kunci.Gol && daftarPemain[j].Assist < kunci.Assist)) {
			daftarPemain[j+1] = daftarPemain[j]
			j--
		}
		daftarPemain[j+1] = kunci
	}

	fmt.Println("\nHasil Sorting :")
	for i := 0; i < n; i++ {
		fmt.Printf("%s %s %d %d\n", daftarPemain[i].NamaDepan, daftarPemain[i].NamaBelakang, daftarPemain[i].Gol, daftarPemain[i].Assist)
	}
}