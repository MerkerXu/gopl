package main

import (
    "fmt"
    "os"
    "strconv"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("usage: %s <number1> <number2>\n", os.Args[0])
    }
    m, err := strconv.Atoi(os.Args[1])
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    n, err := strconv.Atoi(os.Args[2])
    if err != nil {
        fmt.Printf("%v\n", err)
    }
    fmt.Printf("gcd(%d,%d) = %d\n", m, n, gcd(m, n))
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }

    return x
}
