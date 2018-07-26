package main

import (
    "os"
    "io"
    "log"
    "net/http"
)

func main() {
    newFile, err := os.Create("kuaidingyue.html")
    if err != nil {
        log.Fatal(err)
    }
    defer newFile.Close()

    url := "http://www.kuaidingyue.com"
    response, err := http.Get(url)
    defer response.Body.Close()

    numBytesWritten, err := io.Copy(newFile, response.Body)
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Downloaded %b byte file.\n", numBytesWritten)
}
