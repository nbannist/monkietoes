package monkietoes

import (
    //"fmt"
    "html/template"
    "net/http"
   	"github.com/gorilla/mux"
)

func init() {
	// 
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/", simpleHello)

	// 
	http.Handle("/", appRouter)
}

var welcome = template.Must(template.ParseFiles(
  "templates/_app-base.template.html",
  "templates/welcome.template.html",
))

func simpleHello(w http.ResponseWriter, req *http.Request) {
    if err := welcome.Execute(w, nil); err != nil {
      http.Error(w, err.Error(), http.StatusInternalServerError)
  }
}

