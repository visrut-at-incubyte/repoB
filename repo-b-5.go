package main

import "fmt"

func main() {
    a, b, c := 3, 7, 5
    max := a
    if b > max {
        max = b
    }
    if c > max {
        max = c
    }
    fmt.Printf("Max: %d\n", max)
}
