package main

import (
	"fmt"
)

func main() {
	var suaraCalon [21]int
	totalSuaraMasuk := 0
	totalSuaraSah := 0

	for {
		var nilai int
		fmt.Scan(&nilai)

		if nilai == 0 {
			break
		}
		totalSuaraMasuk++

		if nilai >= 1 && nilai <= 20 {
			totalSuaraSah++
			suaraCalon[nilai]++
		}
	}

	fmt.Printf("Suara masuk: %d\n", totalSuaraMasuk)
	fmt.Printf("Suara sah: %d\n", totalSuaraSah)

	for i := 1; i <= 20; i++ {
		if suaraCalon[i] > 0 {
			fmt.Printf("%d : %d\n", i, suaraCalon[i])
		}
	}
}