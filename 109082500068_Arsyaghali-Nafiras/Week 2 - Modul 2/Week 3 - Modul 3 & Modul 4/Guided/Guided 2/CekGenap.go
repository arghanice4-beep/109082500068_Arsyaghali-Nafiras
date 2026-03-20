package main

import "fmt"

func cekGenap(angka int) bool {
    if angka%2 == 0 {
        return true
    }

    return false
}

func main() {
    angkaTes := 8

    hasilGenap := cekGenap(angkaTes)
    fmt.Printf("Apakah %d itu bilangan genap? %t\n", angkaTes, hasilGenap)
}