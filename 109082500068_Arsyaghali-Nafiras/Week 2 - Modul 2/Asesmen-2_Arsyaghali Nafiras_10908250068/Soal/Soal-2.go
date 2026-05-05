package main

import "fmt"

const nMax int = 51

type mahasiswa struct {
	NIM   string
	nama  string
	nilai int
}

type arrayMahasiswa [nMax]mahasiswa

func cariNilaiPertama(T arrayMahasiswa, n int, nimCari string) int {
	for i := 0; i < n; i++ {
		if T[i].NIM == nimCari {
			return T[i].nilai
		}
	}
	return -1 
}

func cariNilaiTerbesar(T arrayMahasiswa, n int, nimCari string) int {
	max := -1
	found := false
	for i := 0; i < n; i++ {
		if T[i].NIM == nimCari {
			if !found || T[i].nilai > max {
				max = T[i].nilai
				found = true
			}
		}
	}
	return max
}

func main() {
	var T arrayMahasiswa
	var n int
	var nimCari string

	fmt.Print("Masukkan jumlah data : ")
	fmt.Scan(&n)

	if n > nMax {
		n = nMax
	}

	for i := 0; i < n; i++ {
		fmt.Printf("Masukkan data ke-%d : ", i+1)
		fmt.Scan(&T[i].NIM, &T[i].nama, &T[i].nilai)
	}

	fmt.Print("Masukkan NIM mahasiswa yang ingin dicari nilai pertama dan nilai terbesarnya : ")
	fmt.Scan(&nimCari)

	nilaiPertama := cariNilaiPertama(T, n, nimCari)
	nilaiTerbesar := cariNilaiTerbesar(T, n, nimCari)

	if nilaiPertama != -1 {
		fmt.Printf("Nilai pertama dari NIM %s adalah %d\n", nimCari, nilaiPertama)
		fmt.Printf("Nilai terbesar dari NIM %s adalah %d\n", nimCari, nilaiTerbesar)
	} else {
		fmt.Printf("Data dengan NIM %s tidak ditemukan.\n", nimCari)
	}
}