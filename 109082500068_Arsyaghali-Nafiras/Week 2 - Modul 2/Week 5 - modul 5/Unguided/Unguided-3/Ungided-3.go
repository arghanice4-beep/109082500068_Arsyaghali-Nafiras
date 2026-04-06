package main
import "fmt"

func tampilkanGanjil(current int, n int) {
	if current > n {
		return
	}

	fmt.Printf("%d ", current)
	tampilkanGanjil(current+2, n)
}

func main() {
	var n int

	fmt.Scan(&n)

	tampilkanGanjil(1, n)
	fmt.Println()
}