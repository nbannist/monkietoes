package main

import (
	//"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

var welcome = template.Must(template.ParseFiles(
	"templates/_app-base.template.html",
	"templates/welcome.template.html",
))

func simpleHello(w http.ResponseWriter, req *http.Request) {
	if err := welcome.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func init() {
	//
	appRouter := mux.NewRouter()
	appRouter.HandleFunc("/", simpleHello)

	//
	http.Handle("/", appRouter)
}
