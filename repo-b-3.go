package main

import "fmt"

func main() {
    n, a, b := 10, 0, 1
    fmt.Print("Fibonacci: ")
    for i := 0; i < n; i++ {
        fmt.Printf("%d ", a)
        a, b = b, a+b
    }
}
