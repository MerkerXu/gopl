package main

import (
	"fmt"
	"gopl/ch2/popcount"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s <integer>\n", os.Args[0])
		os.Exit(1)
	}
	input, err := strconv.ParseUint(os.Args[1], 10, 64)
	if err != nil {
		fmt.Printf("%v\n", err)
	}
	fmt.Printf("popcount(%d)=%d\n", input, popcount.PopCount(input))
}
