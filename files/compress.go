package main

import (
    "os"
    "compress/gzip"
    "log"
)

func main() {
    outputFile, err := os.Create("test.txt.gz")
    if err != nil {
        log.Fatal(err)
    }

    gzipWriter := gzip.NewWriter(outputFile)
    defer gzipWriter.Close()

    _, err = gzipWriter.Write([]byte("Gophers rule!\n"))
    if err != nil {
        log.Fatal(err)
    }

    log.Println("Compressed data written to file.")
}
