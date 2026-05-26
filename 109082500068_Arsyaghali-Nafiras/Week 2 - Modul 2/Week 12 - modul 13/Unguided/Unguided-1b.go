package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var m int
		fmt.Scan(&m)

		var ganjil []int
		var genap []int

		for j := 0; j < m; j++ {
			var nomorRumah int
			fmt.Scan(&nomorRumah)

			if nomorRumah%2 != 0 {
				ganjil = append(ganjil, nomorRumah)
			} else {
				genap = append(genap, nomorRumah)
			}
		}

		for x := 0; x < len(ganjil); x++ {
			for y := x + 1; y < len(ganjil); y++ {
				if ganjil[x] > ganjil[y] {
					ganjil[x], ganjil[y] = ganjil[y], ganjil[x] // Tukar posisi
				}
			}
		}

		for x := 0; x < len(genap); x++ {
			for y := x + 1; y < len(genap); y++ {
				if genap[x] < genap[y] {
					genap[x], genap[y] = genap[y], genap[x] 
				}
			}
		}

		hasil := append(ganjil, genap...)

		for k, val := range hasil {
			if k > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}