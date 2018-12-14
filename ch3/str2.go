package main

import (
    "fmt"
    "reflect"
    "strconv"
    "unsafe"
)

type stringInterner map[string]string

func (si stringInterner) Intern(s string) string {
    if interned, ok := si[s]; ok {
        return interned
    }
    si[s] = s
    return s
}

func stringptr(s string) uintptr {
    return (*reflect.StringHeader)(unsafe.Pointer(&s)).Data
}

func main() {
    si := stringInterner{}
    s1 := si.Intern("12")
    s2 := si.Intern(strconv.Itoa(12))
    fmt.Println(stringptr(s1) == stringptr(s2))
}
