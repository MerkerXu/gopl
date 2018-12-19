package main

import (
    "fmt"
    "bytes"
    "strings"
    "math"
)

func comma(s string) string {
    var buf bytes.Buffer
    dotIndex := strings.Index(s, ".")
    intPart := s[:dotIndex]
    floatPart := s[dotIndex+1:]
    leftLen := len(intPart)%3
    intParts := int(math.Ceil(float64(len(intPart))/3.0))
    for i:=0; i<intParts; i++ {
        var start, end int
        if i==0 {
            start = 0
            if leftLen == 0 {
                leftLen = 3
            }
            end = start + leftLen
        } else {
            start = leftLen+(i-1)*3
            end = start + 3
        }
        buf.Write([]byte(intPart[start:end]))
        if intParts>1 && i<intParts-1 {
            buf.WriteString(",")
        }
    }
    buf.WriteString(".")
    buf.Write([]byte(floatPart))
    return buf.String()
}

func main() {
    fmt.Println(comma(".56789123"))
    fmt.Println(comma("12.56789123"))
    fmt.Println(comma("123.56789123"))
    fmt.Println(comma("12345.56789123"))
    fmt.Println(comma("12."))
}
