package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	for *n < NMAX {
		fmt.Scanf("%c", &input)
		if input == '.' {
			break
		}
		if input != ' ' && input != '\n' {
			t[*n] = input
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikanArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i] // Simple swap di Go
	}
}

func palindrom(t tabel, n int) bool {
	original := t
	balikanArray(&t, n)
	for i := 0; i < n; i++ {
		if original[i] != t[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	isPal := palindrom(tab, m)

	balikanArray(&tab, m)
	fmt.Print("Reverse: ")
	cetakArray(tab, m)

	fmt.Println("Palindrom:", isPal)
}
