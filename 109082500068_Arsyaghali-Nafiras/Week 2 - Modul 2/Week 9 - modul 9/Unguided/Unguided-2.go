package main

import (
	"fmt"
	"math"
)

func main() {
	var n, x, delIdx, target int

	// Input N dan isi array
	fmt.Print("Masukkan jumlah N: ")
	fmt.Scan(&n)
	arr := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	// a. Menampilkan keseluruhan array dengan cara paling singkat
	fmt.Println("\na. Semua elemen:", arr)

	// b. Indeks ganjil
	fmt.Print("b. Indeks ganjil: ")
	for i := 1; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}

	// c. Indeks genap (mulai dari 0, lompat 2)
	fmt.Print("\nc. Indeks genap: ")
	for i := 0; i < len(arr); i += 2 {
		fmt.Print(arr[i], " ")
	}

	// d. Kelipatan x (langsung melompat sebanyak x)
	fmt.Print("\nd. Masukkan x: ")
	fmt.Scan(&x)
	fmt.Print("   Kelipatan x: ")
	if x > 0 {
		for i := 0; i < len(arr); i += x {
			fmt.Print(arr[i], " ")
		}
	}

	// e. Menghapus indeks
	fmt.Print("\ne. Hapus indeks ke-: ")
	fmt.Scan(&delIdx)
	arr = append(arr[:delIdx], arr[delIdx+1:]...) // Rumus jitu Golang untuk hapus elemen
	fmt.Println("   Array setelah dihapus:", arr)

	// f & g. Rata-rata dan Standar Deviasi (menggunakan range agar lebih rapi)
	var sum, sumSquares float64
	for _, val := range arr {
		sum += float64(val)
	}
	mean := sum / float64(len(arr))

	for _, val := range arr {
		sumSquares += math.Pow(float64(val)-mean, 2)
	}
	stdDev := math.Sqrt(sumSquares / float64(len(arr)))

	fmt.Printf("f. Rata-rata: %.2f\n", mean)
	fmt.Printf("g. Standar Deviasi: %.2f\n", stdDev)

	// h. Frekuensi angka
	fmt.Print("h. Cari frekuensi angka: ")
	fmt.Scan(&target)
	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	fmt.Printf("   Angka %d muncul %d kali\n", target, count)
}
