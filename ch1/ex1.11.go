package main

import (
    "fmt"
    "io"
    "io/ioutil"
    "net/http"
    "os"
    "time"
    "strings"
)

func main() {
    start := time.Now()
    ch := make(chan string)
    data, err := ioutil.ReadFile("inputurls.txt")

    outFile, err := os.OpenFile("outfile.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
    if err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
    defer outFile.Close()

    urls := strings.Split(string(data), "\n")
    for _, url := range  urls {
        if !strings.HasPrefix(url, "http://") {
            url = "http://" + url
        }
        go fetch(url, ch)
    }
    for range urls {
        outFile.WriteString(<-ch + "\n")
    }
    outFile.WriteString(fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds()))
}

func fetch(url string, ch chan<- string) {
    start := time.Now()
    resp, err := http.Get(url)
    if err != nil {
        ch <- fmt.Sprint(err)
        ch <- url
        return
    }
    nbytes, err := io.Copy(ioutil.Discard, resp.Body)
    resp.Body.Close()
    if err != nil {
        ch <- fmt.Sprintf("while reading %s: %v", url, err)
        return
    }
    secs := time.Since(start).Seconds()
    ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
