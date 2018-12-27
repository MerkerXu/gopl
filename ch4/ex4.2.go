package main

import (
    "fmt"
    "crypto/sha256"
    "crypto/sha512"
    "flag"
    "strings"
)

func main() {
    hashPtr := flag.String("hash", "SHA256", "Hash Algorithm to use")
    flag.Parse()

    inputString := strings.Join(flag.Args(), " ")

    if *hashPtr == "SHA384" {
        fmt.Printf("%x\n", sha512.Sum384([]byte(inputString)))
    } else if *hashPtr == "SHA256" {
        fmt.Printf("%x\n", sha256.Sum256([]byte(inputString)))
    } else if *hashPtr == "SHA512" {
        fmt.Printf("%x\n", sha512.Sum512([]byte(inputString)))
    } else {
        panic("algorithm not exists: " + *hashPtr)
    }
}
