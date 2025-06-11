package main

import (
    "encoding/json"
    "net/http"
    "runtime"
)

type Metrics struct {
    OS     string
    CPUs   int
    Memory uint64
}

func metricsHandler(w http.ResponseWriter, r *http.Request) {
    m := Metrics{
        OS:     runtime.GOOS,
        CPUs:   runtime.NumCPU(),
        Memory: 16 * 1024 * 1024 * 1024,
    }
    json.NewEncoder(w).Encode(m)
}

func main() {
    http.HandleFunc("/metrics", metricsHandler)
    http.ListenAndServe(":9000", nil)
}
