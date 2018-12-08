package main

import (
    "fmt"
    "os"
    "strconv"
)

func gcd(num1, num2 uint64) uint64 {
    for num2!=0 {
        num1, num2 = num2, num1%num2
    }
    return num1
}

func main() {
    if len(os.Args) != 3 {
        fmt.Printf("usage: %s <num1> <num2>\n", os.Args[0])
        os.Exit(1)
    }
    num1, err1 := strconv.ParseUint(os.Args[1], 10, 64)
    num2, err2 := strconv.ParseUint(os.Args[2], 10, 64)
    if err1!=nil || err2 != nil {
        fmt.Printf("usage: %s <num1> <num2>\n", os.Args[0])
        os.Exit(1)
    }
    fmt.Printf("gcd(%d,%d)=%d\n", num1, num2, gcd(num1, num2))
}
