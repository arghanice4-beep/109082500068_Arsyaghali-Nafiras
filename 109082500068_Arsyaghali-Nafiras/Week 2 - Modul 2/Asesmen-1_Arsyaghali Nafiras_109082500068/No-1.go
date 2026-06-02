package main
import "fmt"

const pi float64 = 3.14

func volume(r, t float64) float64 {
	return pi * r * r * t
}

func massa(r, t, p float64) float64 {
	return volume(r, t) * p
}

func display(m1, m2 float64) {
	if m1 == m2 {
		fmt.Println("BALANCE")
	} else {
		var selisih float64
		if m1 > m2 {
			selisih = m1 - m2
		} else {
			selisih = m2 - m1
		}
		fmt.Printf("Selisih massa zat cair kiri dan massa zat cair kanan : %v\n", selisih)
	}
}

func main() {
	var r float64 // jari-jari
	var tKiri, tKanan float64 //tinggi zat cair kiri dan kanan
	var mjKiri, mjKanan float64 //massa jenis zat cair
	var massaKiri, massaKanan float64 // massa zat cair kiri dan kanan

	// masukkan jari-jari alas tabung
	fmt.Print("Masukkan jari-jari alas tabung : ")
	fmt.Scan(&r)
	fmt.Println() 

	// masukkan tinggi zat cair di tabung kiri, beserta massa jenisnya
	fmt.Print("Masukkan tinggi zat cair tabung kiri : ")
	fmt.Scan(&tKiri)
	fmt.Print("Masukkan massa jenis zat cair tabung kiri : ")
	fmt.Scan(&mjKiri)
	fmt.Println()

	// masukkan tinggi zat cair di tabung kanan, beserta massa jenisnya
	fmt.Print("Masukkan tinggi zat cair tabung kanan : ")
	fmt.Scan(&tKanan)
	fmt.Print("Masukkan massa jenis zat cair tabung kanan : ")
	fmt.Scan(&mjKanan)
	fmt.Println()

	// hitung massa zat cair di tabung kiri dan kanan
	massaKiri = massa(r, tKiri, mjKiri)
	massaKanan = massa(r, tKanan, mjKanan)

	// tampilkan hasil dari proses penimbangan
	display(massaKiri, massaKanan)
}