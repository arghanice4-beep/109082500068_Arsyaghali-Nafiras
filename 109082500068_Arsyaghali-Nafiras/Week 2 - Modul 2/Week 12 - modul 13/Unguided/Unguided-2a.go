package main

import "fmt"

func insertionSort(arr []int) {
	n := len(arr)
	for i := 1; i < n; i++ {
		key := arr[i]
		j := i - 1

		for j >= 0 && arr[j] > key {
			arr[j+1] = arr[j]
			j = j - 1
		}
		arr[j+1] = key
	}
}

func main() {
	var arr []int
	var input int

	for {
		fmt.Scan(&input)
		if input < 0 {
			break
		}
		arr = append(arr, input)
	}

	if len(arr) == 0 {
		return
	}
	insertionSort(arr)

	for i, val := range arr {
		fmt.Print(val)
		if i < len(arr)-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()

	if len(arr) < 2 {
		fmt.Println("Data berjarak 0")
		return
	}

	selisihAwal := arr[1] - arr[0]
	isTetap := true

	for i := 1; i < len(arr)-1; i++ {
		selisihSekarang := arr[i+1] - arr[i]
		if selisihSekarang != selisihAwal {
			isTetap = false
			break
		}
	}

	if isTetap {
		fmt.Printf("Data berjarak %d\n", selisihAwal)
	} else {
		fmt.Println("Data berjarak tidak tetap")
	}
}