package main

import (
    "fmt"
    "crypto/sha256"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func PopCount(ptr *[32]byte) int {
    count := 0
    for i:=0; i<32; i++ {
        count += int(pc[ptr[i]])
    }

    return count
}

func main() {
    sum1 := sha256.Sum256([]byte("x"))
    sum2 := sha256.Sum256([]byte("X"))
    var sum [32]uint8;

    for i:=0; i<len(sum1); i++ {
        sum[i] = sum1[i] ^ sum2[i]
    }
    //fmt.Printf("%x, %[1]T\n", sum)
    fmt.Println(PopCount(&sum))
}
