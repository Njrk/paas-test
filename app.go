package main

import (
    "fmt"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprint(w, "<h2> Test app3 (go) </h2>")
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }

    err := http.ListenAndServe(":" + port, nil)
    if err != nil {
        fmt.Println(err)
    }
}
