package main

import (
    "fmt"
    "io"
    "net/http"
    "os"
    "time"
    "strings"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    var filename string
    filecount := make(map[string]int)

    for _, url := range os.Args[1:] {
        if !strings.HasPrefix(url, "http://") {
            filecount[url] += 1
            filename = url + fmt.Sprintf(".%d", filecount[url]) + ".txt"
            url = "http://" + url
        } else {
            url = url[7:]
            filecount[url] += 1
            filename = url + fmt.Sprintf(".%d", filecount[url]) +  ".txt"
        }

        go fetch(url, ch, filename)
    }
    for range os.Args[1:] {
        fmt.Println(<-ch)
    }
    fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string, filename string) {
    start := time.Now()

    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }
    savedFile,err := os.Create(filename)
    if err != nil {
        ch <- fmt.Sprint(err)
        return
    }

    nbytes, err := io.Copy(savedFile, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
