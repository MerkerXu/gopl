package main

import (
    "fmt"
)

type Weekday int

func main() {

    const (
        Sunday Weekday = iota
        Monday
        Tuesday
        Wednesday
        Thursday
        Friday
        Saturday
    )

    fmt.Printf("%T %[1]v\n", Monday)

}
