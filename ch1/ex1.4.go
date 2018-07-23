package main

import (
    "fmt"
    "os"
    "io/ioutil"
    "strings"
)

func main() {
    counts := make(map[string]int)
    filenames := make(map[string]string)

    for _, filename := range os.Args[1:] {
        data, err := ioutil.ReadFile(filename)
        if err != nil {
            fmt.Fprintf(os.Stderr, "ex1.4: %v\n", err)
            continue
        }
        for _, line := range strings.Split(string(data), "\n") {
            counts[line]++
            filenames[line] += " " + filename
        }
    }

    for line, n := range counts {
        if n > 1 {
            fmt.Printf("%d\t%s\t%s\n", n, line, filenames[line])
        }
    }
}
