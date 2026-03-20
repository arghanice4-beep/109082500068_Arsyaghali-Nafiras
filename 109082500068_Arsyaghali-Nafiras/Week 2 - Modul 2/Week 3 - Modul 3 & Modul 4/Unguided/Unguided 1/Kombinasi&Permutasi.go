package main

import "fmt"

func factorial(n int) int {
	if n == 0 || n == 1 {
		return 1
	}
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

func permutasi(n, r int) int {
	return factorial(n) / factorial(n-r)
}

func kombinasi(n, r int) int {
	return factorial(n) / (factorial(r) * factorial(n-r))
}

func main() {
	var a, b, c, d int

	fmt.Print("Masukkan nilai a: ")
	fmt.Scan(&a)
	fmt.Print("Masukkan nilai b: ")
	fmt.Scan(&b)
	fmt.Print("Masukkan nilai c: ")
	fmt.Scan(&c)
	fmt.Print("Masukkan nilai d: ")
	fmt.Scan(&d)

	if a < c || b < d {
		fmt.Println("Syarat a >= c dan b >= d tidak terpenuhi.")
		return
	}

	fmt.Printf("hasil Permutasi %d dan %d: %d\n", a, c, permutasi(a, c))

	fmt.Printf("hasil Kombinasi %d dan %d: %d\n", a, c, kombinasi(a, c))

	fmt.Printf("hasil Permutasi %d dan %d: %d\n", b, d, permutasi(b, d))

	fmt.Printf("hasil Kombinasi %d dan %d: %d\n", b, d, kombinasi(b, d))
}
