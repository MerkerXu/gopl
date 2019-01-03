package main

import "fmt"

const LEN int = 10

func reverse(ptr *[LEN]int) {
    s := *ptr
    for i, j := 0, len(s)-1; i<j; i,j = i+1,j-1 {
        s[i], s[j] = s[j], s[i]
    }
}

func main() {
    a := [LEN]int{0,1,2,3,4,5}
    fmt.Println(a)
    reverse(&a)
    fmt.Println(a)
}
