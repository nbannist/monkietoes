package appmod

import (
	"github.com/gorilla/mux"
	"github.com/nbannist/monkietoes/frontend"
	"log"
	"net/http"
)

var Cfg Config

func LoadStandaloneRoutes() {
	/*
		This function is only needed for the main.go part of the app
	*/
	log.Print("Loading Routes for Standalone Server...")
}

func init() {

	Cfg = Config{
		Env:      DEV,
		Platform: "appengine", // assumption
	}
	//
	log.Print("Creating Router...\n")
	appRouter := mux.NewRouter()
	//
	appRouter.HandleFunc("/", frontend.Handlers["root"])
	log.Print("MONKIETOES.INIT(): root route attached")
	//
	docRouter := appRouter.PathPrefix("/d/").Subrouter()
	// DOCUMENT ROUTS
	docRouter.HandleFunc("/", frontend.Handlers["getNewDoc"])
	log.Print("MONKIETOES.INIT(): newDoc route attached")
	//
	docRouter.HandleFunc("/{docid}", frontend.Handlers["getDoc"]).Methods("GET")
	log.Print("MONKIETOES.INIT(): getDoc route attached")
	//
	docRouter.HandleFunc("/{docid}", frontend.Handlers["saveDoc"]).Methods("POST")
	log.Print("MONKIETOES.INIT(): saveDoc route attached")
	//
	//appRouter.HandleFunc("/d/new", frontend.Handlers["createDoc"])
	//log.Print("MONKIETOES.INIT(): saveDoc route attached")
	//
	//appRouter.HandleFunc("/d/", frontend.Handlers["saveDoc"])
	//log.Print("MONKIETOES.INIT(): saveDoc route attached")

	// add router
	http.Handle("/", appRouter)

	//
	// Logging
	//
	log.Print("Router created.\n")
	log.Print("init() complete...\n")
	log.Print("config platform = " + Cfg.Platform)
}
