package main

import (
	"fmt"
	"strings"
)
func main() {
	target := [4]string{"merah", "kuning", "hijau", "ungu"}
	Percobaan := true
	for i := 1; i <= 5; i++ {
		var g1, g2, g3, g4 string
		fmt.Printf("Percobaan %d: ", i)
		_, err := fmt.Scanln(&g1, &g2, &g3, &g4)
		if err != nil {
			Percobaan = false
			break
		}
		inputUser := [4]string{
			strings.ToLower(g1), 
			strings.ToLower(g2), 
			strings.ToLower(g3), 
			strings.ToLower(g4),
		}
		if inputUser != target {
			Percobaan = false
		}
	}
	fmt.Printf("Berhasil: %t\n", Percobaan)
}