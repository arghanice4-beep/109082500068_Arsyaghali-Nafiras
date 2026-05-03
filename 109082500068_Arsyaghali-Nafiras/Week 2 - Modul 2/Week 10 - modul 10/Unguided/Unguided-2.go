package main

import (
	"fmt"
)

func main() {
	var beratIkan [1000]float64
	var x, y int

	fmt.Print("Masukkan jumlah ikan (x) dan kapasitas wadah (y): ")
	fmt.Scan(&x, &y)

	fmt.Print("Masukkan berat masing-masing ikan: ")
	for i := 0; i < x; i++ {
		fmt.Scan(&beratIkan[i])
	}

	jumlahWadah := (x + y - 1) / y
	totalBeratPerWadah := make([]float64, jumlahWadah)

	var totalSeluruhBerat float64 = 0

	idxWadah := 0
	for i := 0; i < x; i++ {
		totalBeratPerWadah[idxWadah] += beratIkan[i]
		
		if (i+1)%y == 0 {
			idxWadah++
		}
	}

	for i := 0; i < jumlahWadah; i++ {
		fmt.Printf("total berat wadah %d: %.2f ", i+1, totalBeratPerWadah[i])
		totalSeluruhBerat += totalBeratPerWadah[i]
	}
	fmt.Println()

	rataRata := totalSeluruhBerat / float64(jumlahWadah)
	fmt.Printf("berat rata-rata ikan per wadah: %.2f\n", rataRata)
}