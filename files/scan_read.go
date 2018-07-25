package main

import (
    "os"
    "log"
    "fmt"
    "bufio"
)

func main() {
    file, err := os.Open("test.txt")
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(file)
    scanner.Split(bufio.ScanWords)

    success := scanner.Scan()
    if success == false {
        err = scanner.Err()
        if err == nil {
            log.Println("Scan completed and reached EOF")
        } else {
            log.Fatal(err)
        }
    }

    fmt.Println("First word found:", scanner.Text())
}
