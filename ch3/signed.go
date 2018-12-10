package main

import "fmt"

func main() {
    medals := []string{"gold", "silver", "bronze"}
    for i:= len(medals)-1; i>=0; i-- {
        fmt.Println(medals[i])
    }

    o := 0666
    fmt.Printf("%d %[1]o %#[1]o\n", o)
    x := int64(0xdeadbeef)
    fmt.Printf("%d %[1]x %#[1]x %#[1]X\n", x)
}
