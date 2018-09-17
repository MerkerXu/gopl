package main

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"
)

func sayhello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("Key:", k)
		fmt.Println("Val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello!"+time.Now().String())
}

func main() {
	http.HandleFunc("/", sayhello)
	err := http.ListenAndServe(":4000", nil)
	if err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
