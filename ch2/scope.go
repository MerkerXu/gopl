package main

import "fmt"
import "os"
import "log"

var cwd string

func init() {
    var err error
    cwd, err = os.Getwd()
    // !! NOTE: short variable decaration
    // cwd, err := os.Getwd()
    if err != nil {
        log.Fatalf("os.Getwd failed: %v", err)
    }
    log.Printf("Working directory = %s", cwd)
}

func main() {
    x := "hello!"
    for i:=0; i<len(x); i++ {
        x := x[i]
        if x != '!' {
            x := x + 'A' - 'a'
            fmt.Printf("%c", x)
        } else {
            fmt.Printf("%c", x)
        }
    }
    log.Printf("Working directory = %s", cwd)
}
