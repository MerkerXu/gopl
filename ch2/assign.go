package main

import "fmt"

func main() {
    fmt.Printf("gcd(%d,%d) = %d\n", 18, 27, gcd(18, 27))
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }

    return x
}
