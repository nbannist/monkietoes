package monkietoes

import (
    "fmt"
    "net/http"
)

func init() {
    http.HandleFunc("/", simpleHello)
}

func simpleHello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

