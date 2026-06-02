package main
import "fmt"

func tanggunganHari(jumlahHari int, tujuan string) int {
	/* Mengembalikan jumlah hari maksimum yang biaya perjalanannya
	   ditanggung oleh Tel-U berdasarkan lama study tour (jumlahHari) dan tujuan
	   (domestik/mancanegara) */
	if tujuan == "domestik" {
		if jumlahHari > 3 {
			return 3
		}
		return jumlahHari
	} else if tujuan == "mancanegara" {
		if jumlahHari > 8 {
			return 8
		}
		return jumlahHari
	}
	return 0
}

func biayaPerHari(jumlahMhs int) int {
	/* Menghitung biaya tour domestik per hari yang ditanggung oleh Tel-U untuk
	   jumlah mahasiswa sebanyak jumlahMhs*/
	//yang dihitung biayaPerHari hanya untuk domestik saja, karena untuk manca negara biayaPerHari nya = biayaPerHari domestik x 1,5(1,5 kali biaya domestik)
	biayaDomestikPerMhs := 70000 + 250000 + 300000
	return jumlahMhs * biayaDomestikPerMhs
}

func perhitunganBiaya(jumlahMhs, lamaPerjalanan int, tujuan string, totalBiaya *float64) {
	// parameter totalBiaya menggunakan method pass by reference

	/* I.S. Terdefinisi jumlah mahasiswa, lama hari study tour, dan tujuan
	   perjalanan (domestic/mancanegara)
	   F.S. Telah dihitung biaya perjalanan yang ditanggung Tel-U */

	// Panggil salah satu fungsi/prosedur untuk menghitung lama perjalanan
	hariDitanggung := tanggunganHari(lamaPerjalanan, tujuan)

	// Panggil fungsi/prosedur untuk menghitung biaya total tour domestic seluruh mahasiswa
	biayaDomestikTotal := biayaPerHari(jumlahMhs)

	// Hitung biaya study tour seluruh mahasiswa jika tujuan domestik atau mancanegara
	if tujuan == "domestik" {
		*totalBiaya = float64(hariDitanggung * biayaDomestikTotal)
	} else if tujuan == "mancanegara" {
		*totalBiaya = float64(hariDitanggung*biayaDomestikTotal) * 1.5
	}
}

func main() {
	var jumlah, lama int
	var tujuan string
	var biaya float64

	// lakukan proses masukan atau input di sini
	fmt.Print("masukkan jumlah mahasiswa : ")
	fmt.Scan(&jumlah)
	fmt.Print("Masukkan lama hari study tour : ")
	fmt.Scan(&lama)
	fmt.Print("Masukkan tujuan study tour (domestik/mancanegara) : ")
	fmt.Scan(&tujuan)

	// hitung biaya perjalanan yang dikeluarkan Tel-U dengan memanggil subprogram yang tepat
	perhitunganBiaya(jumlah, lama, tujuan, &biaya)

	// tampilkan biaya
	fmt.Println()
	fmt.Printf("Biaya perjalanan yang harus dikeluarkan Tel-U : Rp. %.0f\n", biaya)
}