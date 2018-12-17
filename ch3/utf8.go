package main

import "fmt"
import "unicode/utf8"

func HasPrefix(s, prefix string) bool {
    return len(s) >= len(prefix) && s[:len(prefix)] == prefix
}

func HasSuffix(s, suffix string) bool {
    return len(s) >= len(suffix) && s[len(s)-len(suffix):] == suffix
}

func Contains(s, substr string) bool {
    for i:=0; i<len(s); i++ {
        if HasPrefix(s[i:], substr) {
            return true
        }
    }
    return false
}

func main() {
    const placeOfInterest = `⌘`

    fmt.Printf("plain string:")
    fmt.Printf("%s", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("quoted string:")
    fmt.Printf("%+q", placeOfInterest)
    fmt.Printf("\n")

    fmt.Printf("hex bytes:")
    for i:=0; i<len(placeOfInterest);i++ {
        fmt.Printf("%x ", placeOfInterest[i])
    }
    fmt.Printf("\n")

    fmt.Println(HasPrefix("abc", "ab"))
    fmt.Println(HasPrefix("abc", "ac"))
    fmt.Println(HasSuffix("abc", "bc"))
    fmt.Println(HasSuffix("abc", "c"))
    fmt.Println(HasSuffix("abc", "ac"))
    fmt.Println(Contains("abc", "ac"))
    fmt.Println(Contains("abc", "bc"))

    s := "Hello, 世界"
    fmt.Println(len(s))
    fmt.Println(utf8.RuneCountInString(s))

    for i:=0; i<len(s); {
        r, size := utf8.DecodeRuneInString(s[i:])
        fmt.Printf("%d\t%c\n", i, r)
        i += size
    }

    for i,r := range "Hello, 世界" {
        fmt.Printf("%d\t%q\t%d\n", i, r, r)
    }
}
