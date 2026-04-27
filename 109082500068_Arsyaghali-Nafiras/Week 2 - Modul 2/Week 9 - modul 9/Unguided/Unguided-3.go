package main

import "fmt"

func main() {
	var a, b string
	var sA, sB int
	var hasil []string

	fmt.Print("Klub A : "); fmt.Scan(&a)
	fmt.Print("Klub B : "); fmt.Scan(&b)

	for i := 1; ; i++ {
		fmt.Printf("Pertandingan %d : ", i)
		fmt.Scan(&sA, &sB)

		if sA < 0 || sB < 0 { break }

		if sA > sB {
			hasil = append(hasil, a)
		} else if sB > sA {
			hasil = append(hasil, b)
		} else {
			hasil = append(hasil, "Draw")
		}
	}

	for i, v := range hasil {
		fmt.Printf("Hasil %d : %s\n", i+1, v)
	}
	fmt.Println("Pertandingan selesai")
}