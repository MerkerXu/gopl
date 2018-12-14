package main

import (
	"fmt"
	"reflect"
	"unsafe"
    "strconv"
)

func stringptr(s string) uintptr {
    return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

func main() {
    s1 := "1234"
    s2 := s1[:2]
    fmt.Println(stringptr(s1) == stringptr(s2))
    fmt.Println(s1 == s2)
    s2 = "12" + "34"
    fmt.Println(stringptr(s1) == stringptr(s2))
    s2 = strconv.Itoa(1234)
    fmt.Println(stringptr(s1) == stringptr(s2))
}
