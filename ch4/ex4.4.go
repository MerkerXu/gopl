package main

import "fmt"

func rotate(a []int, k int) {
}

func main() {
    a := [...]int{1,2,3,4,5,6}
    fmt.Println(a)
    rotate(a[:], 2)
    fmt.Println(a)
}
