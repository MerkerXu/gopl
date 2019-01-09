package main

import "fmt"

func rotate(a []int, k int) {
	length := len(a)
	k = k % length
	count := 0
	for startPos := 0; count < length; startPos++ {
		nextPos := (startPos + (length - k)) % length
		nextVal := a[nextPos]
		a[nextPos] = a[startPos]
		currPos := nextPos
		count++
		for startPos != currPos {
			nextPos := (currPos + (length - k)) % length
			nextVal, a[nextPos] = a[nextPos], nextVal
			count++
			currPos = nextPos
		}
	}
}

func main() {
	a := [...]int{1, 2, 3, 4, 5, 6, 8}
	fmt.Println(a)
	rotate(a[:], 4)
	fmt.Println(a)
}
