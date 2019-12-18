package main

import (
    "fmt"
    "log"
    "net/http"
    "io_text"
)

func main() {
    http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
        fmt.Fprintf(w, io_text.ReadTxt("/dcgm_metrics.txt"))
    })

    log.Fatal(http.ListenAndServe(":8080", nil))
}