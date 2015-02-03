package monkietoes

import (
    "fmt"
    "net/http"
   	"github.com/gorilla/mux"
)

func init() {
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/", simpleHello)


	http.Handle("/", appRouter)
}

func simpleHello(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintln(w, "Hello, World!")
}

