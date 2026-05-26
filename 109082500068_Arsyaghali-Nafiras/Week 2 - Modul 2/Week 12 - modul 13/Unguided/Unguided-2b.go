package main

import "fmt"

const nMax = 7919

type Buku struct {
	id        string
	judul     string
	penulis   string
	penerbit  string
	eksemplar int
	tahun     int
	rating    int
}

type DaftarBuku [nMax]Buku

func DaftarkanBuku(pustaka *DaftarBuku, n *int) {
	fmt.Scan(n)
	for i := 0; i < *n; i++ {
		fmt.Scan(
			&pustaka[i].id, 
			&pustaka[i].judul, 
			&pustaka[i].penulis, 
			&pustaka[i].penerbit, 
			&pustaka[i].eksemplar, 
			&pustaka[i].tahun, 
			&pustaka[i].rating,
		)
	}
}

func CetakTerfavorit(pustaka DaftarBuku, n int) {
	if n > 0 {
		maxRating := pustaka[0].rating
		idxMax := 0
		for i := 1; i < n; i++ {
			if pustaka[i].rating > maxRating {
				maxRating = pustaka[i].rating
				idxMax = i
			}
		}
		fmt.Printf("%s %s %s %d\n", pustaka[idxMax].judul, pustaka[idxMax].penulis, pustaka[idxMax].penerbit, pustaka[idxMax].tahun)
	}
}

func UrutBuku(pustaka *DaftarBuku, n int) {
	for i := 1; i < n; i++ {
		temp := pustaka[i]
		j := i - 1

		for j >= 0 && pustaka[j].rating < temp.rating {
			pustaka[j+1] = pustaka[j]
			j--
		}
		pustaka[j+1] = temp
	}
}

func Cetak5Terbaru(pustaka DaftarBuku, n int) {
	batas := 5
	if n < 5 {
		batas = n
	}

	for i := 0; i < batas; i++ {
		fmt.Println(pustaka[i].judul)
	}
}

func CariBuku(pustaka DaftarBuku, n int, r int) {
	kiri := 0
	kanan := n - 1
	ketemu := false
	tengah := 0

	for kiri <= kanan && !ketemu {
		tengah = (kiri + kanan) / 2

		if pustaka[tengah].rating == r {
			ketemu = true
		} else if pustaka[tengah].rating < r {
			kanan = tengah - 1
		} else {
			kiri = tengah + 1
		}
	}

	if ketemu {
		fmt.Printf("%s %s %s %d %d %d\n", 
			pustaka[tengah].judul, 
			pustaka[tengah].penulis, 
			pustaka[tengah].penerbit, 
			pustaka[tengah].tahun, 
			pustaka[tengah].eksemplar, 
			pustaka[tengah].rating,
		)
	} else {
		fmt.Println("Tidak ada buku dengan rating seperti itu")
	}
}

func main() {
	var Pustaka DaftarBuku
	var nPustaka int
	var rCari int

	DaftarkanBuku(&Pustaka, &nPustaka)
	CetakTerfavorit(Pustaka, nPustaka)
	UrutBuku(&Pustaka, nPustaka)
	Cetak5Terbaru(Pustaka, nPustaka)
	fmt.Scan(&rCari)
	CariBuku(Pustaka, nPustaka, rCari)
}