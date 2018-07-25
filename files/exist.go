package main

import (
    "log"
    "os"
)

var (
    fileInfo *os.FileInfo
    err error
)

func main() {
    fileInfo, err := os.Stat("test.txt")
    if err != nil {
        if os.IsNotExist(err) {
            log.Fatal("File does not exist.")
        }
    }
    log.Println("File does exist. File infomation:")
    log.Println(fileInfo)
}
