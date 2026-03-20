package main
import "fmt"
func main() {
	var tahun int
	var kabisat bool
	fmt.Print("Masukkan tahun: ")
	fmt.Scanln(&tahun)
	kabisat = (tahun%400 == 0) || (tahun%4 == 0 && tahun%100 != 0)
	fmt.Printf("kabisat: %t\n", kabisat)
}