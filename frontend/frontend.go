package frontend

import (
	"fmt"
	"github.com/gorilla/mux"
	"html/template"
	"log"
	"net/http"
)

type DocData struct {
	Docid string
	Title string
	Body  string
	//...
}

var Handlers map[string]func(http.ResponseWriter, *http.Request)

// STATIC PAGES

func webRoot(w http.ResponseWriter, r *http.Request) {
	log.Print("webRoot Handler")
	var rootTemplate = template.Must(template.ParseFiles(
		"templates/_app-base.template.html",
		"templates/welcome.template.html",
	))

	// execute template; capture error;
	if err := rootTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func aboutSite(w http.ResponseWriter, r *http.Request) {
	log.Print("aboutSite Handler")
	var rootTemplate = template.Must(template.ParseFiles(
		"templates/_app-base.template.html",
		"templates/welcome.template.html",
	))

	// execute template; capture error;
	if err := rootTemplate.Execute(w, nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

// DYNAMIC PAGES

func getDocument(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	data := DocData{
		Docid: vars["docid"],
		Title: "TITLE - " + vars["docid"],
		Body:  "temp body",
	}

	var getDocumentTemplate = template.Must(template.ParseFiles(
		"templates/_app-base.template.html",
		"templates/document.get.template.html",
		"templates/footer-header.template.html",
		"templates/footer-skirt.template.html",
	))

	log.Print("getDocument Handler")
	log.Print("docid: " + vars["docid"])
	fmt.Println("GET DOCUMENT")

	// execute template; capture error;
	if err := getDocumentTemplate.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
func getNewDocument(w http.ResponseWriter, r *http.Request) {

	log.Print("newDocument Handler")
	fmt.Println("NEW DOCUMENT")
}
func saveDocument(w http.ResponseWriter, r *http.Request) {

	log.Print("saveDocument Handler")
	fmt.Println("SAVE DOCUMENT")
}
func deleteDocument(w http.ResponseWriter, r *http.Request) {

	log.Print("deleteDocument Handler")
	fmt.Println("DELETE DOCUMENT")
}

func init() {
	// maps take a key value in brakets [] then afterward the return type
	// if it's a function the signature of the function is required
	// INIT
	Handlers = make(map[string]func(http.ResponseWriter, *http.Request))

	Handlers["root"] = webRoot
	Handlers["getNewDoc"] = getNewDocument
	Handlers["getDoc"] = getDocument
	Handlers["saveDoc"] = saveDocument
	Handlers["deleteDoc"] = deleteDocument
}
