package main

import (
    "os"
    "io/ioutil"
    "log"
    "fmt"
)

func main() {
    tempDirPath, err := ioutil.TempDir("", "myTempDir")
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println("Temp dir created:", tempDirPath)

    tempFile, err := ioutil.TempFile(tempDirPath, "myTempFile.txt")
    if err  != nil {
        log.Fatal(err)
    }
    fmt.Println("Temp file created:", tempFile.Name())

    err = tempFile.Close()
    if err != nil {
        log.Fatal(err)
    }

    err = os.Remove(tempFile.Name())
    if err != nil {
        log.Fatal(err)
    }

    err = os.Remove(tempDirPath)
    if err != nil {
        log.Fatal(err)
    }
}
