package main
import "fmt"

func cetakDeret(n int) {
	for n != 1 {
		fmt.Printf("%d ", n) 
		
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
	}
	fmt.Printf("%d\n", n)
}

func main() {
	var n int

	_, err := fmt.Scan(&n)
	
	if err != nil {
		return
	}
	if n > 0 && n < 1000000 {
		cetakDeret(n)
	} else {
		fmt.Println("Masukan tidak valid. Harus bilangan positif < 1000000")
	}
}