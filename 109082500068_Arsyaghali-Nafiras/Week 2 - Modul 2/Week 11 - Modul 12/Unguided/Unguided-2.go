package main

import "fmt"

type Calon struct {
	Nomor int
	Suara int
}

func main() {
	hitungSuara := make([]int, 21)

	suaraMasuk := 0
	suaraSah := 0

	for {
		var nilai int
		fmt.Scan(&nilai)

		if nilai == 0 {
			break
		}
		suaraMasuk++

		if nilai >= 1 && nilai <= 20 {
			suaraSah++
			hitungSuara[nilai]++
		}
	}

	var daftarCalon []Calon
	for i := 1; i <= 20; i++ {
		daftarCalon = append(daftarCalon, Calon{Nomor: i, Suara: hitungSuara[i]})
	}

	idxKetua := 0
	for i := 1; i < len(daftarCalon); i++ {
		if daftarCalon[i].Suara > daftarCalon[idxKetua].Suara {
			idxKetua = i
		} else if daftarCalon[i].Suara == daftarCalon[idxKetua].Suara {
			if daftarCalon[i].Nomor < daftarCalon[idxKetua].Nomor {
				idxKetua = i
			}
		}
	}
	ketua := daftarCalon[idxKetua]

	idxWakil := -1
	for i := 0; i < len(daftarCalon); i++ {
		if i == idxKetua {
			continue
		}
		if idxWakil == -1 {
			idxWakil = i
			continue
		}
		if daftarCalon[i].Suara > daftarCalon[idxWakil].Suara {
			idxWakil = i
		} else if daftarCalon[i].Suara == daftarCalon[idxWakil].Suara {
			if daftarCalon[i].Nomor < daftarCalon[idxWakil].Nomor {
				idxWakil = i
			}
		}
	}
	wakil := daftarCalon[idxWakil]

	fmt.Printf("Suara masuk: %d\n", suaraMasuk)
	fmt.Printf("Suara sah: %d\n", suaraSah)
	fmt.Printf("Ketua RT: %d\n", ketua.Nomor)
	fmt.Printf("Wakil ketua: %d\n", wakil.Nomor)
}