package main

import (
    "log"
    "net/http"
    "fmt"
    "html"
)

func index(w http.ResponseWriter, r *http.Request) {
    log.Println("Showing index")
    fmt.Fprintf(w, fmt.Sprintf("Hello. you requested %s", html.EscapeString(r.URL.Path)))
}

func main() {
    log.Println("Booting...")
    http.HandleFunc("/", index)
    log.Fatal(http.ListenAndServe(":80", nil))
}


