package main
import "fmt"

type arrdata [5]string

func seqSearch(arr arrdata, binatangCari string) int {
	var found int = -1
	for i:=0; i<len(arr); i++ {
		if arr[i] == binatangCari {
			found = i
			break
		}
	}
	return found
}

func main() {
	var arrbinatang arrdata

	for i:=0; i<len(arrbinatang); i++ {
		fmt.Printf("Masukkan data binatang indeks ke-%d: ", i)
		fmt.Scanln(&arrbinatang[i])
	}

	var binatangCari string
	fmt.Print("Masukkan nama binatang yang mau dicari: ")
	fmt.Scanln(&binatangCari)

	var idxcari int 
	idxcari = seqSearch(arrbinatang, binatangCari)
	
	if idxcari != -1 {
		fmt.Printf("Binatang '%s' ditemukan di indeks %d\n", binatangCari, idxcari)
	} else {
		fmt.Printf("Binatang '%s' tidak ditemukan\n", binatangCari)
	}
}