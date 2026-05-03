package main

import (
	"fmt"
)

func main() {
	var beratKelinci [1000]float64
	var n int

	fmt.Print("Masukkan jumlah kelinci (N): ")
	fmt.Scan(&n)

	if n > 1000 {
		n = 1000
	}

	for i := 0; i < n; i++ {
		fmt.Print("Masukkan berat kelinci ke-", i+1, ": ")
		fmt.Scan(&beratKelinci[i])
	}

	min := beratKelinci[0]
	max := beratKelinci[0]

	for i := 1; i < n; i++ {
		if beratKelinci[i] < min {
			min = beratKelinci[i]
		}
		if beratKelinci[i] > max {
			max = beratKelinci[i]
		}
	}

	fmt.Printf("Berat terkecil: %.2f, Berat terbesar: %.2f\n", min, max)
}