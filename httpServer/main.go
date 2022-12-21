package main

import (
    "fmt"
    "os"
    "net/http"
)

func main() {
    http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request) {
        for _, e := range os.Environ() {
            fmt.Fprintf(w, "%s\n", e)
        }
    })

    http.ListenAndServe(":8080",nil)
}