package main

import "fmt"

func basename(s string) string {
    // discard last '/' and everything before
    for i:= len(s)-1; i>=0; i-- {
        if s[i] == '/' {
            fmt.Printf("%d:%d\n", len(s), i)
            s = s[i+1:]
            break
        }
    }
    // preserve everything before last '.'
    for i:=len(s)-1; i>=0; i-- {
        if s[i] == '.' {
            s = s[:i]
            break
        }
    }

    return s
}

func main() {
    fmt.Println(basename("a/b/c.go/"))
    fmt.Println(basename("a/b/c.go"))
    fmt.Println(basename("a.d.go"))
    fmt.Println(basename("abc"))
}
