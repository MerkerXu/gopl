package main

import (
    "fmt"
    "os"
)

func main() {
    if len(os.Args) != 2 {
        fmt.Printf("usage: %s <filename>", os.Args[0])
        os.Exit(1)
    }
    err := os.Truncate(os.Args[1], 0)
    if err != nil {
        fmt.Fprintf(os.Stderr, "%v\n", err)
    }
}
