package main

import (
	"fmt"
	"math"
)

func main() {
	var cx1, cy1, r1 float64
	var cx2, cy2, r2 float64
	var x, y float64

	// Input data
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	// Hitung jarak titik ke pusat lingkaran 1 & 2
	jarak1 := math.Sqrt(math.Pow(x-cx1, 2) + math.Pow(y-cy1, 2))
	jarak2 := math.Sqrt(math.Pow(x-cx2, 2) + math.Pow(y-cy2, 2))

	// Cek kondisi
	inL1 := jarak1 <= r1
	inL2 := jarak2 <= r2

	// Output sesuai syarat
	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
