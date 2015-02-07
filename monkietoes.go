package monkietoes

import (
	//"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

// For Non-AppEngine Environments
func main() {
	http.HandleFunc("/", hello)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

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
