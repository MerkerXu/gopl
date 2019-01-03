package main

import (
	"fmt"
)

func appendInt(x []int, y int) []int {
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        // there is a room to grow. Extend the slice
        z = x[:zlen]
    } else {
        // there is insufficient space. Allocate a new array.
        // Grow by doubling, for amortized linear complexity
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)  // a built-in function
    }
    z[len(x)] = y
    
    return z
}

func main() {
	var runes []rune
	for _, r := range "Hello, 世界" {
		runes = append(runes, r)
	}

	fmt.Printf("%q\n", runes)

    var x, y []int
    for i:=0; i<12; i++ {
        y = appendInt(x, i)
        fmt.Printf("%d\tlen=%d\t\tcap=%d\t\t%v\n", i, len(y), cap(y), y)
        x = y
    }

    var  m []int
    m = append(m, 1)
    m = append(m, 2, 3)
    m = append(m, 4, 5, 6)
    m = append(m, m...)
    fmt.Println(m)
}
