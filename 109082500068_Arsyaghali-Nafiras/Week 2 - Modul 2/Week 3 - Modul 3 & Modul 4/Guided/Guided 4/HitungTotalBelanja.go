package main

import "fmt"

func hitungtotalbelanja(harga int, jumlah int){
    var total int
    total = harga * jumlah
    fmt.Println("Total belanja: ", total)

    hitungdiskon(total)

}

func hitungdiskon(total int){
    var diskon, hargaakhir int
    diskon =total * 10 / 100
    hargaakhir = total - diskon

    fmt.Println("Diskon 10%: ", diskon)
    fmt.Println("Harga akhir (setelah diskon): ", hargaakhir)
}

func main() {
    var harga, jumlah int

    fmt.Print("Masukkan harga barang: ")
    fmt.Scan(&harga)
    fmt.Print("Masukkan jumlah barang: ")
    fmt.Scan(&jumlah)

    hitungtotalbelanja(harga, jumlah)
}