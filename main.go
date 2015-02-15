// +build !appengine

package main

import (
	"github.com/nbannist/monkietoes/appmod"
	"log"
	"net/http"
)

// Server for Non-AppEngine Environments
func main() {
	appmod.LoadStandaloneRoutes()
	appmod.Cfg.Platform = "standalone"

	log.Print("Loading main()...")
	log.Print("Listening on port: 8070")
	log.Print("config platform = " + appmod.Cfg.Platform)

	if err := http.ListenAndServe(":8070", nil); err != nil {
		log.Print("Error... :/ \n")
		panic(err)
	}

	log.Print("main() END")
}
