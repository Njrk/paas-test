package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        fmt.Fprint(w, "<h2> Test app3 (go) </h2>")
	log.Printf("Request from %s: %s %s\n", r.RemoteAddr, r.Method, r.URL)
    })

    port := os.Getenv("PORT")
    if port == "" {
        port = "5000"
    }

    log.Printf("Server is listening on :%s\n", port)
    err := http.ListenAndServe(":"+port, nil)
    if err != nil {
        log.Fatalf("Error starting the server: %v", err)
    }
}
