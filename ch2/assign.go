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

    fmt.Printf("fib(%d)=%d\n", m, fib(m))
}

func gcd(x, y int) int {
    for y != 0 {
        x, y = y, x%y
    }

    return x
}

func fib(n int) int {
    x, y := 0, 1
    for i:=0; i<n; i++ {
        x, y = y, x+y
    }
    return x
}
