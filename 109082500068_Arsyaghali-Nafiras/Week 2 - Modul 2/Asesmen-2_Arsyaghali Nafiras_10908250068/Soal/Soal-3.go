package main

import "fmt"

const nProv = 10

type NamaProv [nProv]string
type PopProv [nProv]int
type TumbuhProv [nProv]float64

func main() {
	var provinsi NamaProv
	var populasi PopProv
	var pertumbuhan TumbuhProv
	var namaCari string

	fmt.Printf("--- Masukkan Nama Provinsi, Populasi Provinsi, Angka Pertumbuhan Provinsi ---\n")
	InputData(&provinsi, &populasi, &pertumbuhan)

	fmt.Print("Masukkan nama provinsi yang dicari: ")
	fmt.Scan(&namaCari)

	idxTercepat := ProvinsiTercepat(pertumbuhan)
	fmt.Printf("\nProvinsi dengan angka pertumbuhan tercepat: %s\n", provinsi[idxTercepat])

	idxCari := IndeksProvinsi(provinsi, namaCari)
	fmt.Printf("Data provinsi yang dicari (%s): %d\n", namaCari, idxCari+1)

	fmt.Println("\n--- Prediksi Jumlah Penduduk Tahun Depan Pada Provinsi Dengan Pertumbuhan Diatas 2% ---")
	Prediksi(provinsi, populasi, pertumbuhan)
}

func InputData(prov *NamaProv, pop *PopProv, tumbuh *TumbuhProv) {
	for i := 0; i < nProv; i++ {
		fmt.Printf("Masukkan data ke-%d: ", i+1)
		fmt.Scan(&prov[i], &pop[i], &tumbuh[i])
	}
}

func ProvinsiTercepat(tumbuh TumbuhProv) int {
	idxMax := 0
	for i := 1; i < nProv; i++ {
		if tumbuh[i] > tumbuh[idxMax] {
			idxMax = i
		}
	}
	return idxMax
}

func IndeksProvinsi(prov NamaProv, nama string) int {
	for i := 0; i < nProv; i++ {
		if prov[i] == nama {
			return i
		}
	}
	return -1
}

func Prediksi(prov NamaProv, pop PopProv, tumbuh TumbuhProv) {
	for i := 0; i < nProv; i++ {
		if tumbuh[i] > 0.02 {
			prediksiPenduduk := float64(pop[i]) * (tumbuh[i] + 1)
			fmt.Printf("%-15s %.0f\n", prov[i], prediksiPenduduk)
		}
	}
}