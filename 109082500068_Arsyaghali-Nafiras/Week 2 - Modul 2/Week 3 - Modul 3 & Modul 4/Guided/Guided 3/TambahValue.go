package main

import "fmt"

func tambahvalue(x int){
    x = x + 10
    fmt.Printf("Nilai x di dalam fungsi tambahvalue (pass by value): %d\n", x)

}

func tambahreference(x *int){
    *x = *x + 10
    fmt.Printf("Nilai x di dalam fungsi tambahreference (pass by reference): %d\n", *x)
}

func main() {
    var y int = 5

    fmt.Println("nilai y awal: ", y)
    tambahvalue(y)
    fmt.Println("nilai x setelah pemanggilan tambahvalue: ", y)
    tambahreference(&y)
    fmt.Println("nilai x setelah pemanggilan tambahreference: ", y)
}