package main

import "fmt"

type suhu float64

func CelciusToReamur(c suhu) suhu {
	return c * 0.8
}

func CelciusToFahrenheit(c suhu) suhu {
	return (c * 9 / 5) + 32
}

func CelciusToKelvin(c suhu) suhu {
	return c + 273.15
}

func main() {
	var input suhu

	fmt.Println("=== KONVERTER CELCIUS ===")
	fmt.Print("Masukkan suhu (celcius) : ")
	fmt.Scan(&input)

	reamur := CelciusToReamur(input)
	fahrenheit := CelciusToFahrenheit(input)
	kelvin := CelciusToKelvin(input)

	fmt.Println("") 
	fmt.Printf("%.0f celcius = %.1f reamur\n", input, reamur)
	fmt.Printf("%.0f celcius = %.1f fahrenheit\n", input, fahrenheit)
	fmt.Printf("%.0f celcius = %.2f kelvin\n", input, kelvin)
}